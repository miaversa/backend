package router

import (
	"github.com/go-chi/chi"
	"net/http"
)

func CartRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", view)
	return r
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cart view"))
}
