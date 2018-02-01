package cart

import (
	"github.com/miaversa/backend/cookie"
	"github.com/miaversa/backend/model"
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type service struct {
	cookie cookie.Service
}

func New(cookie cookie.Service) *service {
	return &service{cookie: cookie}
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.view(w, r)
		return
	}
	s.update(w, r)
}

func (s *service) view(w http.ResponseWriter, r *http.Request) {
	c, err := s.cookie.GetCart(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t := template.New("cart_view.html")
	t.Parse(string(templates.MustAsset("cart_view.html")))
	err = t.Execute(w, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *service) update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	method := r.PostFormValue("_method")
	if strings.ToLower(method) == "delete" {
		s.RemoveItem(w, r, r.PostFormValue("sku"))
		return
	}
	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i := model.CartItem{
		Product: model.Product{
			SKU:   r.PostFormValue("sku"),
			Name:  r.PostFormValue("name"),
			Price: price,
		},
		Quantity: 1,
	}
	s.AddItem(w, r, i)
}

func (s *service) AddItem(w http.ResponseWriter, r *http.Request, i model.CartItem) {
	c, err := s.cookie.GetCart(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.AddItem(i)
	err = s.cookie.SaveCart(w, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/carrinho", http.StatusFound)
}

func (s *service) RemoveItem(w http.ResponseWriter, r *http.Request, sku string) {
	c, err := s.cookie.GetCart(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.RemoveItem(sku)
	err = s.cookie.SaveCart(w, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/carrinho", http.StatusFound)
}
