package cart

import (
	"github.com/miaversa/backend/model"
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Path for the routing
var Path string = "/carrinho"

// CartStore defines the api for the cart storage
type CartStore interface {
	GetCart(r *http.Request) (model.Cart, error)
	SaveCart(w http.ResponseWriter, c model.Cart) error
	DropCart(w http.ResponseWriter)
}

type handler struct {
	store CartStore
}

// New creates a new Cart Handler
func New(store CartStore) *handler {
	return &handler{store: store}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := h.view(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := h.update(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) view(w http.ResponseWriter, r *http.Request) (err error) {
	if c, err := h.store.GetCart(r); err == nil {
		t := template.New("cart.html")
		t.Parse(string(templates.MustAsset("cart.html")))
		err = t.Execute(w, c)
	}
	return
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()
	method := r.PostFormValue("_method")
	if strings.ToLower(method) == "delete" {
		index, err := strconv.Atoi(r.PostFormValue("index"))
		if err != nil {
			return err
		}
		return h.removeItem(w, r, index)
	}
	if price, err := strconv.ParseFloat(r.PostFormValue("price"), 64); err == nil {
		p := model.Product{
			SKU:     r.PostFormValue("sku"),
			Name:    r.PostFormValue("name"),
			Price:   price,
			Options: []model.ProductOption{},
		}

		optSize := r.PostFormValue("option_size")
		if "" != optSize {
			p.Options = append(p.Options, model.ProductOption{Name: "size", Value: optSize})
		}
		err = h.addItem(w, r, p)
	}
	return
}

func (h *handler) addItem(w http.ResponseWriter, r *http.Request, p model.Product) (err error) {
	if c, err := h.store.GetCart(r); err == nil {
		c.AddProduct(p)
		if err = h.store.SaveCart(w, c); err == nil {
			http.Redirect(w, r, Path, http.StatusFound)
		}
	}
	return
}

func (h *handler) removeItem(w http.ResponseWriter, r *http.Request, index int) (err error) {
	if c, err := h.store.GetCart(r); err == nil {
		c.RemoveProduct(index)
		if err = h.store.SaveCart(w, c); err == nil {
			http.Redirect(w, r, Path, http.StatusFound)
		}
	}
	return
}
