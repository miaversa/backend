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
	cartHandler := handler.NewCartHandler(cartStorage)

	r.Get("/carrinho", handler.HandlerError(cartHandler.GetCart))
	r.Post("/carrinho", handler.HandlerError(cartHandler.Update))

	http.ListenAndServe(":8080", r)
}
