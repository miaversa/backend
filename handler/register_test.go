package handler_test

/*
import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/mem"
	"github.com/miaversa/backend/register"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHandler_View(t *testing.T) {
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := register.New(sessionService, customerService)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}
	// TODO: check response body
}

func TestHandler_View_With_Session(t *testing.T) {
	sessionService := mem.NewSessionService("session")
	sessionService.Email = "maria@gmail.com"
	customerService := mem.NewCustomerService()
	handler := register.New(sessionService, customerService)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("Received non-302 response: %d\n", rr.Code)
	}

	location, ok := rr.HeaderMap["Location"]
	if !ok {
		t.Fatal("location not found in the header")
	}
	locationString := location[0]
	if locationString != register.DefaultRedirectPath {
		t.Fatal("location mismatch")
	}
}

func TestHandler_Register_Invalid(t *testing.T) {
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := register.New(sessionService, customerService)

	form := url.Values{}
	form.Add("email", "")

	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}

	// TODO: check response body
}

func TestHandler_Register_Valid(t *testing.T) {
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := register.New(sessionService, customerService)

	email := "maria@gmail.com"
	name := "Maria Madalena"
	password := "123456"
	form := url.Values{}
	form.Add("email", email)
	form.Add("name", name)
	form.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("Received non-302 response: %d\n", rr.Code)
	}

	hc := rr.Header().Get("Set-Cookie")
	if hc != "session=maria@gmail.com" {
		t.Fatal("cookie error")
	}

	// TODO: check response body
}

func TestHandler_Double_Register(t *testing.T) {
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := register.New(sessionService, customerService)

	email := "maria@gmail.com"
	name := "Maria Madalena"
	password := "123456"

	customerService.Put(customer.Customer{
		Email:    email,
		Name:     name,
		Password: password,
	})

	form := url.Values{}
	form.Add("email", email)
	form.Add("name", name)
	form.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}

	hc := rr.Header().Get("Set-Cookie")
	if hc != "" {
		t.Fatal("cookie error")
	}

	// TODO: check response body
}
*/
