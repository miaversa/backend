package handler_test

import (
	"github.com/miaversa/backend/handler"
	"github.com/miaversa/backend/mem"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHandler_View(t *testing.T) {
	store := mem.NewCartStorage()
	handler := handler.NewCartHandler(store)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	handler.GetCart(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}
	// TODO: check response body
}

func TestHandler_Delete_Invalid(t *testing.T) {
	store := mem.NewCartStorage()
	handler := handler.NewCartHandler(store)

	form := url.Values{}
	form.Add("_method", "delete")
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	c, _ := store.GetCart()
	_, err := handler.DeleteProduct(rr, req, c)
	if err == nil {
		t.Fatal("esperava um erro")
	}
	form.Add("index", "asd")
	req, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()
	c, _ = store.GetCart()
	_, err = handler.DeleteProduct(rr, req, c)
	if err == nil {
		t.Fatal("esperava um erro")
	}

}

func TestHandler_Delete_Valid(t *testing.T) {
	store := mem.NewCartStorage()
	handler := handler.NewCartHandler(store)

	form := url.Values{}
	form.Add("_method", "delete")
	form.Add("index", "0")
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	c, _ := store.GetCart()
	_, err := handler.DeleteProduct(rr, req, c)

	if err != nil {
		t.Fatal(err)
	}
}

func TestHandler_Add_Valid(t *testing.T) {
	store := mem.NewCartStorage()
	handler := handler.NewCartHandler(store)

	sku := "x"
	form := url.Values{}
	form.Add("sku", sku)
	form.Add("name", "x")
	form.Add("price", "10.5")
	form.Add("option_size", "15")
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	c, _ := store.GetCart()
	_, err := handler.AddProduct(rr, req, c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandler_Add_Invalid_Price(t *testing.T) {
	store := mem.NewCartStorage()
	handler := handler.NewCartHandler(store)
	sku := "x"
	price := "1xdsf"
	form := url.Values{}
	form.Add("sku", sku)
	form.Add("name", "x")
	form.Add("price", price)
	form.Add("option_size", "15")
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	c, _ := store.GetCart()
	_, err := handler.AddProduct(rr, req, c)
	if err == nil {
		t.Fatal("esperava um erro de conversão")
	}
}

func TestHandler_Update_Delete(t *testing.T) {
	store := mem.NewCartStorage()
	handler := handler.NewCartHandler(store)
	form := url.Values{}
	form.Add("_method", "delete")
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler.Update(rr, req)
}
