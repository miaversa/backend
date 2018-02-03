package cart

import (
	"github.com/miaversa/backend/model"
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var Path string = "/carrinho"

type CartStore interface {
	GetCart(r *http.Request) (model.Cart, error)
	SaveCart(w http.ResponseWriter, c model.Cart) error
	DropCart(w http.ResponseWriter)
}

type handler struct {
	store CartStore
}

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
		return h.removeItem(w, r, r.PostFormValue("sku"))
	}
	if price, err := strconv.ParseFloat(r.PostFormValue("price"), 64); err == nil {
		i := model.CartItem{
			Product: model.Product{
				SKU:     r.PostFormValue("sku"),
				Name:    r.PostFormValue("name"),
				Price:   price,
				Options: []model.ProductOption{},
			},
			Quantity: 1,
		}
		optSize := r.PostFormValue("option_size")
		if "" != optSize {
			i.Product.Options = append(i.Product.Options, model.ProductOption{Name: "size", Value: optSize})
		}
		err = h.addItem(w, r, i)
	}
	return
}

func (h *handler) addItem(w http.ResponseWriter, r *http.Request, i model.CartItem) (err error) {
	if c, err := h.store.GetCart(r); err == nil {
		c.AddItem(i)
		if err = h.store.SaveCart(w, c); err == nil {
			http.Redirect(w, r, Path, http.StatusFound)
		}
	}
	return
}

func (h *handler) removeItem(w http.ResponseWriter, r *http.Request, sku string) (err error) {
	if c, err := h.store.GetCart(r); err == nil {
		c.RemoveItem(sku)
		if err = h.store.SaveCart(w, c); err == nil {
			http.Redirect(w, r, Path, http.StatusFound)
		}
	}
	return
}