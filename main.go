package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/miaversa/backend/assets"
	"github.com/miaversa/backend/cart"
	"github.com/miaversa/backend/config"
	"github.com/miaversa/backend/cookie"
	"github.com/miaversa/backend/login"
	"github.com/miaversa/backend/payment"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	config.Load()

	r := chi.NewRouter()
	r.Use(middleware.NoCache)
	r.Use(middleware.StripSlashes)

	secure := true
	if viper.GetBool("debug") {
		secure = false
	}

	hashKey := viper.GetString("cookie.hashKey")
	blockKey := viper.GetString("cookie.blockKey")

	sessionService := cookie.NewSessionService(viper.GetString("cookie.session.name"), hashKey, blockKey, secure)
	loginService := login.New(sessionService)

	cartStore := cookie.NewCartStore(viper.GetString("cookie.cart.name"), hashKey, blockKey, secure)
	cartService := cart.New(cartStore)
	assetService := assets.New()
	paymentService := payment.New()

	r.Handle(login.Path, loginService)
	r.Handle(cart.Path, cartService)
	r.Handle(payment.Path, paymentService)
	r.Handle(assets.Path, assetService)
	http.ListenAndServe(":8080", r)
}
