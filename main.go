package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/miaversa/backend/assets"
	"github.com/miaversa/backend/cart"
	"github.com/miaversa/backend/config"
	"github.com/miaversa/backend/cookie"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	config.Load()

	r := chi.NewRouter()
	r.Use(middleware.NoCache)

	secure := true
	if viper.GetBool("debug") {
		secure = false
	}

	hashKey := viper.GetString("cookie.hashKey")
	blockKey := viper.GetString("cookie.blockKey")

	cookieService := cookie.New(viper.GetString("cookie.name"), hashKey, blockKey, secure)
	cartService := cart.New(cookieService)
	assetService := assets.New()

	r.Handle(cart.Path, cartService)
	r.Handle("/assets/{filename}", assetService)
	http.ListenAndServe(":8080", r)
}
