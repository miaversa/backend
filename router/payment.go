package router

import (
	"github.com/go-chi/chi"
	"net/http"
)

func PaymentRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", paymentView)
	return r
}

func paymentView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("payment view"))
}
