package handler_test

import (
	"errors"
	"github.com/miaversa/backend/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func returnNil(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func returnErr(w http.ResponseWriter, r *http.Request) error {
	return errors.New("erro")
}

func TestErrorHandler(t *testing.T) {
	h := handler.HandlerError(returnNil)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatal("200 expected")
	}
	h = handler.HandlerError(returnErr)
	req, _ = http.NewRequest(http.MethodGet, "/", nil)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusInternalServerError {
		t.Fatal("500 expected")
	}
}
