package handler

import (
	"github.com/miaversa/backend/cart"
	"github.com/miaversa/backend/product"
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
	"strconv"
)

type cartHandler struct {
	cartStorage cart.CartStorage
}

func NewCartHandler(cartStorage cart.CartStorage) *cartHandler {
	return &cartHandler{cartStorage}
}

func (h *cartHandler) GetCart(w http.ResponseWriter, r *http.Request) error {
	c, err := h.cartStorage.GetCart()
	if err != nil {
		return err
	}
	t := template.New("cart.html")
	t.Parse(string(templates.MustAsset("cart.html")))
	return t.Execute(w, c)
}

func (h *cartHandler) Update(w http.ResponseWriter, r *http.Request) error {
	c, err := h.cartStorage.GetCart()
	if err != nil {
		return err
	}
	if r.PostFormValue("_method") == "delete" {
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

func (h *cartHandler) AddProduct(w http.ResponseWriter, r *http.Request, c cart.Cart) (cart.Cart, error) {
	p, err := extractProduct(r)
	if err != nil {
		return c, err
	}
	c.AddProduct(p)
	return c, nil

}

func (h *cartHandler) DeleteProduct(w http.ResponseWriter, r *http.Request, c cart.Cart) (cart.Cart, error) {
	index, err := strconv.Atoi(r.PostFormValue("index"))
	if err != nil {
		return c, err
	}
	c.RemoveProduct(index)
	return c, nil
}

func extractProduct(r *http.Request) (product.Product, error) {
	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		return product.Product{}, err
	}
	p := product.Product{
		SKU:     r.PostFormValue("sku"),
		Name:    r.PostFormValue("name"),
		Price:   price,
		Options: []product.Option{},
	}
	optSize := r.PostFormValue("option_size")
	if "" != optSize {
		p.Options = append(p.Options, product.Option{Name: "size", Value: optSize})
	}

	return p, nil
}
