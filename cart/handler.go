package cart

import (
	"github.com/miaversa/backend/model"
	"github.com/miaversa/backend/templates"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Path for the routing
var Path string = "/carrinho"

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
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := h.update(w, r); err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) view(w http.ResponseWriter, r *http.Request) error {
	c, err := h.store.GetCart()
	if err != nil {
		return err
	}
	t := template.New("cart.html")
	t.Parse(string(templates.MustAsset("cart.html")))
	err = t.Execute(w, c)
	return nil
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) error {
	method := r.PostFormValue("_method")
	if strings.ToLower(method) == "delete" {
		index, err := strconv.Atoi(r.PostFormValue("index"))
		if err != nil {
			return err
		}
		return h.removeItem(w, r, index)
	}
	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		return err
	}
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
	return h.addItem(w, r, p)
}

func (h *handler) addItem(w http.ResponseWriter, r *http.Request, p model.Product) error {
	c, err := h.store.GetCart()
	if err != nil {
		return err
	}
	c.AddProduct(p)
	if err = h.store.SaveCart(c); err == nil {
		http.Redirect(w, r, Path, http.StatusFound)
	}
	return nil
}

func (h *handler) removeItem(w http.ResponseWriter, r *http.Request, index int) error {
	c, err := h.store.GetCart()
	if err == nil {
		c.RemoveProduct(index)
		if err = h.store.SaveCart(c); err == nil {
			http.Redirect(w, r, Path, http.StatusFound)
		}
	}
	return err
}
