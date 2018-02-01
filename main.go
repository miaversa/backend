package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/miaversa/backend/assets"
	"github.com/miaversa/backend/cart"
	"github.com/miaversa/backend/cookie"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.NoCache)

	c := cookie.New()

	r.Handle("/carrinho", cart.New(c))
	r.Handle("/assets/{filename}", assets.New())
	http.ListenAndServe(":8080", r)
}
