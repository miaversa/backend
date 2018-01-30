package router

import (
	"github.com/go-chi/chi"
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
)

func CartRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", cartView)
	return r
}

func cartView(w http.ResponseWriter, r *http.Request) {

	cart := Cart{
		Shipping: 123.1,
		Items: []CartItem{
			CartItem{Name: "Anel", Price: 123.10},
			CartItem{Name: "Colar", Price: 127.11},
		},
	}

	t := template.New("cart_view.html")
	t.Parse(string(templates.MustAsset("cart_view.html")))

	err := t.Execute(w, cart)
	if err != nil {
		panic(err)
	}
}

type Cart struct {
	Shipping float64
	Items    []CartItem
}

type CartItem struct {
	Name  string
	Price float64
}

func (c Cart) Total() float64 {
	var sum float64 = 0
	for _, i := range c.Items {
		sum += i.Price
	}

	sum += c.Shipping

	return sum
}
