package router

import (
	"github.com/go-chi/chi"
	"net/http"
)

func CustomerRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", customerView)
	return r
}

func customerView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("customer view"))
}
