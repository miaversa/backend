package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/miaversa/backend/handler"
	"github.com/miaversa/backend/mem"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.NoCache)
	r.Use(middleware.StripSlashes)

	cartStorage := mem.NewCartStorage()
	sessionStorage := mem.NewSessionStorage()
	customerStorage := mem.NewCustomerStorage()

	rh := handler.NewRedirectHandler(sessionStorage)

	cartHandler := handler.NewCartHandler(cartStorage)
	paymentHandler := handler.NewPaymentHandler(sessionStorage, customerStorage)
	loginHandler := handler.NewLoginHandler(sessionStorage, customerStorage)
	registerHandler := handler.NewRegisterHandler(sessionStorage, customerStorage)

	r.Get("/carrinho", handler.HandlerError(cartHandler.GetCart))
	r.Post("/carrinho", handler.HandlerError(cartHandler.Update))
	r.Get("/pagar", handler.HandlerError(paymentHandler.View))
	r.Get("/login", handler.HandlerError(rh.RedirectIfSession(loginHandler.View)))
	r.Post("/login", handler.HandlerError(rh.RedirectIfSession(loginHandler.Auth)))
	r.Get("/cadastro", handler.HandlerError(rh.RedirectIfSession(registerHandler.View)))

	http.ListenAndServe(":8000", r)
}
