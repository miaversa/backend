package payment

import (
	"net/http"
)

var Path string = "/pagar"

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pagar"))
}
