package handler_test

import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/handler"
	"github.com/miaversa/backend/mem"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLogin_View(t *testing.T) {
	sessionStorage := mem.NewSessionStorage()
	customerStorage := mem.NewCustomerStorage()
	handler := handler.NewLoginHandler(sessionStorage, customerStorage)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	err := handler.View(rr, req)
	if err != nil {
		t.Fatal(err)
	}

	if rr.Code != http.StatusOK {
		t.Fatal("esperado codigo 200")
	}
}

func TestLogin_View_With_Session_Redirect(t *testing.T) {
	sessionStorage := mem.NewSessionStorage()
	sessionStorage.Set("maria@gmail.com")
	customerStorage := mem.NewCustomerStorage()
	handler := handler.NewLoginHandler(sessionStorage, customerStorage)
	req, _ := http.NewRequest(http.MethodGet, "/login?redirect=/pagar", nil)
	rr := httptest.NewRecorder()
	err := handler.View(rr, req)
	if err != nil {
		t.Fatal(err)
	}
	if rr.Code != http.StatusFound {
		t.Fatal("esperado codigo 302")
	}
	if rr.Header().Get("Location") != "/pagar" {
		t.Fatal("redirect invalido")
	}
}

func TestLogin_Auth_With_Session(t *testing.T) {
	sessionStorage := mem.NewSessionStorage()
	sessionStorage.Set("maria@gmail.com")
	customerStorage := mem.NewCustomerStorage()
	handler := handler.NewLoginHandler(sessionStorage, customerStorage)
	req, _ := http.NewRequest(http.MethodGet, "/login?redirect=/pagar", nil)
	rr := httptest.NewRecorder()
	err := handler.Auth(rr, req)
	if err != nil {
		t.Fatal(err)
	}
	if rr.Code != http.StatusFound {
		t.Fatal("esperado codigo 302")
	}
	if rr.Header().Get("Location") != "/pagar" {
		t.Fatal("redirect invalido")
	}
}
func TestLogin_Auth_With_Form_Error(t *testing.T) {
	sessionStorage := mem.NewSessionStorage()
	customerStorage := mem.NewCustomerStorage()
	handler := handler.NewLoginHandler(sessionStorage, customerStorage)
	req, _ := http.NewRequest(http.MethodGet, "/login?redirect=/pagar", nil)
	rr := httptest.NewRecorder()
	err := handler.Auth(rr, req)
	if err != nil {
		t.Fatal(err)
	}
	if rr.Code != http.StatusOK {
		t.Fatal("esperado codigo 200")
	}
}

func TestLogin_Auth_With_Login_OK(t *testing.T) {
	sessionStorage := mem.NewSessionStorage()
	customerStorage := mem.NewCustomerStorage()
	customerStorage.Put(customer.Customer{Email: "maria@gmail.com", Name: "Maria", Password: "123456"})
	handler := handler.NewLoginHandler(sessionStorage, customerStorage)
	form := url.Values{}
	form.Add("email", "maria@gmail.com")
	form.Add("password", "123456")
	req, _ := http.NewRequest(http.MethodPost, "/login?redirect=/pagar", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	err := handler.Auth(rr, req)
	if err != nil {
		t.Fatal(err)
	}
	if rr.Code != http.StatusFound {
		t.Fatal("esperado redirect")
	}
	if rr.Header().Get("Location") != "/pagar" {
		t.Fatal("redirect error")
	}
}
func TestLogin_Auth_With_Login_Invalid_Password(t *testing.T) {
	sessionStorage := mem.NewSessionStorage()
	customerStorage := mem.NewCustomerStorage()
	customerStorage.Put(customer.Customer{Email: "maria@gmail.com", Name: "Maria", Password: "123456"})
	handler := handler.NewLoginHandler(sessionStorage, customerStorage)
	form := url.Values{}
	form.Add("email", "maria@gmail.com")
	form.Add("password", "123456x")
	req, _ := http.NewRequest(http.MethodPost, "/login?redirect=/pagar", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	err := handler.Auth(rr, req)
	if err != nil {
		t.Fatal(err)
	}
	if rr.Code != http.StatusOK {
		t.Fatal("esperado codigo 200")
	}
}
