package middleware_test

import (
	"github.com/miaversa/backend/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

type x struct {
	CartID string
}

func (h *x) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetCartID(r.Context())
	if id == "" {
		panic("esperado cartid")
	}
	h.CartID = id
}

func TestCartIDMiddleware(t *testing.T) {
	tx := &x{}
	m := middleware.CartID(tx)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	m.ServeHTTP(rr, req)
	if rr.Header().Get("Set-Cookie") == "" {
		t.Fatal("esperado cookie")
	}
	cookie := &http.Cookie{Name: "cid", Value: "uuid"}
	req.AddCookie(cookie)
	m.ServeHTTP(rr, req)
	if rr.Header().Get("Set-Cookie") == "" {
		t.Fatal("esperado cookie")
	}
	if tx.CartID != "uuid" {
		t.Fatal("erro no cookie passado")
	}
}
