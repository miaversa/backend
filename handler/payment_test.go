package handler_test

import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/handler"
	"github.com/miaversa/backend/mem"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPaymentHandler_View_Without_Session(t *testing.T) {
	sessionStorage := mem.NewSessionStorage()
	customerStorage := mem.NewCustomerStorage()
	handler := handler.NewPaymentHandler(sessionStorage, customerStorage)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	handler.View(rr, req)
	if rr.Code != http.StatusFound {
		t.Fatalf("Received non-302 response: %d\n", rr.Code)
	}
	if rr.Header().Get("Location") != "/login?redirect=/pagar" {
		t.Fatal("redirect invalido")
	}
}

func TestPaymentHandler_View_Without_Shipping(t *testing.T) {
	sessionStorage := mem.NewSessionStorage()
	sessionStorage.Set("maria@gmail.com")
	customerStorage := mem.NewCustomerStorage()
	handler := handler.NewPaymentHandler(sessionStorage, customerStorage)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	handler.View(rr, req)
	if rr.Code != http.StatusFound {
		t.Fatalf("Received non-302 response: %d\n", rr.Code)
	}
	if rr.Header().Get("Location") != "/envio?redirect=/pagar" {
		t.Fatal("redirect invalido")
	}
}

func TestPaymentHandler_View_OK(t *testing.T) {
	sessionStorage := mem.NewSessionStorage()
	sessionStorage.Set("maria@gmail.com")
	customerStorage := mem.NewCustomerStorage()
	customerStorage.SetShippingAddress("maria@gmail.com", customer.ShippingAddress{Street: "rua maria"})
	handler := handler.NewPaymentHandler(sessionStorage, customerStorage)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	handler.View(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}
	if rr.Header().Get("Location") != "" {
		t.Fatal("redirect invalido")
	}
}
