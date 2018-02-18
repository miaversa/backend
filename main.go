package main

import (
	"log"
	"net/http"

	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/pagseguro"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/miaversa/backend/app"
	"github.com/miaversa/backend/config"
	"github.com/miaversa/backend/mem"
)

func main() {
	config.Load()

	cust := &customer.Customer{
		Email: "maria@gmail.com",
		Name:  "Maria",
		ShippingAddress: customer.ShippingAddress{
			City:       "Sao Paulo",
			Complement: "entrada 50",
			Country:    "BRA",
			District:   "Asa Norte",
			Number:     "50",
			PostalCode: "70770000",
			State:      "DF",
			Street:     "SCRN 716",
		},
	}

	cartStorage := mem.NewCartStorage()

	customerStorage := mem.NewCustomerStorage()
	customerStorage.PutCustomer(cust)

	pagSeguroService := pagseguro.New()

	app := app.New(cartStorage, customerStorage, pagSeguroService)

	r := chi.NewRouter()
	r.Use(middleware.NoCache)
	r.Use(middleware.StripSlashes)

	r.Get("/assets", app.Assets)

	r.Get("/carrinho", app.GetCart)
	r.Post("/carrinho", app.UpdateCart)
	r.Get("/pagar", app.ViewPayment)

	r.Get("/pagseguro/session", app.PagSeguroSessionID)

	log.Fatal(http.ListenAndServe(config.Port, r))
}
