package handler

import (
	"github.com/miaversa/backend/model"
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type cartHandler struct {
	cartStorage CartStorage
}

func NewCartHandler(cartStorage CartStorage) *cartHandler {
	return &cartHandler{cartStorage}
}

func (h *cartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	c, err := h.cartStorage.GetCart()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		t := template.New("cart.html")
		t.Parse(string(templates.MustAsset("cart.html")))
		err = t.Execute(w, c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *cartHandler) Update(w http.ResponseWriter, r *http.Request) error {
	c, err := h.cartStorage.GetCart()
	if err != nil {
		return err
	}
	method := r.PostFormValue("_method")
	if strings.ToLower(method) == "delete" {
		c, err = h.DeleteProduct(w, r, c)
		if err != nil {
			return err
		}
	} else {
		c, err = h.AddProduct(w, r, c)
		if err != nil {
			return err
		}
	}
	err = h.cartStorage.SaveCart(c)
	if err != nil {
		return err
	}
	http.Redirect(w, r, "/carrinho", http.StatusFound)
	return nil
}

func (h *cartHandler) AddProduct(w http.ResponseWriter, r *http.Request, c model.Cart) (model.Cart, error) {
	p, err := extractProduct(r)
	if err != nil {
		return c, err
	}
	c.AddProduct(p)
	return c, nil

}

func (h *cartHandler) DeleteProduct(w http.ResponseWriter, r *http.Request, c model.Cart) (model.Cart, error) {
	index, err := strconv.Atoi(r.PostFormValue("index"))
	if err != nil {
		return c, err
	}
	c.RemoveProduct(index)
	return c, nil
}

func extractProduct(r *http.Request) (model.Product, error) {
	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		return model.Product{}, err
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

	return p, nil
}
