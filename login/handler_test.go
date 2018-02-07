package login_test

import (
	"github.com/miaversa/backend/login"
	"github.com/miaversa/backend/mem"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHandler_View(t *testing.T) {
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := login.New(sessionService, customerService)

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
	handler := login.New(sessionService, customerService)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}

	location, ok := rr.HeaderMap["Location"]
	if !ok {
		t.Fatal("location not found in the header")
	}

	locationString := location[0]

	if locationString != login.DefaultRedirectPath {
		t.Fatal("location mismatch")
	}
}

func TestHandler_View_With_Redirect(t *testing.T) {
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := login.New(sessionService, customerService)

	req, err := http.NewRequest(http.MethodGet, "/?redirect=/perfil", nil)
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

func TestHandler_View_With_Session_With_Redirect(t *testing.T) {
	newLocation := "/pagar"
	sessionService := mem.NewSessionService("session")
	sessionService.Email = "maria@gmail.com"
	customerService := mem.NewCustomerService()
	handler := login.New(sessionService, customerService)

	req, err := http.NewRequest(http.MethodGet, "/?redirect="+newLocation, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}

	location, ok := rr.HeaderMap["Location"]
	if !ok {
		t.Fatal("location not found in the header")
	}

	locationString := location[0]

	if locationString != newLocation {
		t.Fatal("location mismatch")
	}
}

func TestHandler_Auth_Invalid(t *testing.T) {
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := login.New(sessionService, customerService)

	req, err := http.NewRequest(http.MethodPost, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}
	// TODO: check response body
}

func TestHandler_Auth_Valid(t *testing.T) {
	email := "maria@gmail.com"
	password := "password"
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := login.New(sessionService, customerService)

	form := url.Values{}
	form.Add("email", email)
	form.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}

	cookie := rr.Header().Get("Set-Cookie")
	if cookie != "session="+email {
		t.Fatal("cookie error")
	}

	location, ok := rr.HeaderMap["Location"]
	if !ok {
		t.Fatal("location not found in the header")
	}
	locationString := location[0]
	if locationString != login.Path+"?redirect="+login.DefaultRedirectPath {
		t.Fatal("location mismatch")
	}
}

func TestHandler_Auth_Valid_No_Auth(t *testing.T) {
	email := "maria@gmail.com"
	password := "password"
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := login.New(sessionService, customerService)

	form := url.Values{}
	form.Add("email", email)
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

	cookie := rr.Header().Get("Set-Cookie")
	if cookie != "" {
		t.Fatal("invalid cookie set")
	}
	// TODO: check body
}

func TestHandler_Auth_Valid_Redirect(t *testing.T) {
	newLocation := "/pagar"
	email := "maria@gmail.com"
	password := "password"
	sessionService := mem.NewSessionService("session")
	customerService := mem.NewCustomerService()
	handler := login.New(sessionService, customerService)

	form := url.Values{}
	form.Add("email", email)
	form.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, "/?redirect="+newLocation, strings.NewReader(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}

	location, ok := rr.HeaderMap["Location"]
	if !ok {
		t.Fatal("location not found in the header")
	}

	locationString := location[0]

	if locationString != newLocation {
		t.Fatal("location mismatch")
	}
}
