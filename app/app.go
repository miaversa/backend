package app

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/miaversa/backend/assets"
	"github.com/miaversa/backend/cart"
	"github.com/miaversa/backend/config"
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/pagseguro"
	"github.com/miaversa/backend/product"
	"github.com/miaversa/backend/templates"
)

type App struct {
	cartStorage      cart.CartStorage
	customerStorage  customer.CustomerStorage
	pagSeguroService pagseguro.PagSeguroService
}

func New(cart cart.CartStorage, customer customer.CustomerStorage, ps pagseguro.PagSeguroService) *App {
	return &App{cart, customer, ps}
}

func (app *App) GetCart(w http.ResponseWriter, r *http.Request) {
	cid := "x"
	c, err := app.cartStorage.GetCart(cid)
	if err != nil && err == cart.CartNotFoundErr {
		c = cart.New("x")
	}
	t := template.New("cart.html")
	t.Parse(string(templates.MustAsset("cart.html")))
	err = t.Execute(w, c)
}

func (app *App) UpdateCart(w http.ResponseWriter, r *http.Request) {
	cid := "x"
	c, err := app.cartStorage.GetCart(cid)
	if err != nil && err == cart.CartNotFoundErr {
		c = cart.New("x")
	}
	if r.PostFormValue("_method") == "delete" {
		index, err := strconv.Atoi(r.PostFormValue("index"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c.RemoveProduct(index)
	} else {
		price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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
		c.AddProduct(p)
	}
	err = app.cartStorage.SaveCart(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/carrinho", http.StatusFound)
}

func (app *App) ViewPayment(w http.ResponseWriter, r *http.Request) {
	sid := "maria@gmail.com"
	ca, err := app.cartStorage.GetCart("x")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cu, err := app.customerStorage.GetCustomer(sid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t := template.New("payment.html")
	t.Parse(string(templates.MustAsset("payment.html")))
	err = t.Execute(w, struct {
		Cart                *cart.Cart
		Customer            *customer.Customer
		PagSeguroJavascript string
	}{ca, cu, config.PagSeguroJavascript})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ServeHTTP sends the file with the correct content type.
func (app *App) Assets(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if strings.HasSuffix(filename, ".css") {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}

	if strings.HasSuffix(filename, ".js") {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	}

	a, err := assets.Asset(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(a)
}

func (app *App) PagSeguroSessionID(w http.ResponseWriter, r *http.Request) {
	id, err := app.pagSeguroService.SessionID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(id))
}
