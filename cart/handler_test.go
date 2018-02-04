package cart_test

import (
	"encoding/base64"
	"github.com/miaversa/backend/cart"
	"github.com/miaversa/backend/mem"
	"github.com/miaversa/backend/model"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHandler_View(t *testing.T) {
	store := mem.NewCartStore("mcart")
	handler := cart.New(store)

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

func TestHandler_Add_Item_Invalid(t *testing.T) {
	store := mem.NewCartStore("mcart")
	handler := cart.New(store)

	req, err := http.NewRequest(http.MethodPost, "/", nil)
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

func TestHandler_Add_Item_Valid(t *testing.T) {

	sku := "asdf"
	name := "name"
	price := "102.3"
	size := "15"

	store := mem.NewCartStore("mcart")
	handler := cart.New(store)

	form := url.Values{}
	form.Add("sku", sku)
	form.Add("name", name)
	form.Add("price", price)
	form.Add("option_size", size)

	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("Received invalid code response: %d\n", rr.Code)
	}

	if rr.HeaderMap.Get("Set-Cookie") != "mcart=eyJzaGlwcGluZyI6MCwiaXRlbXMiOlt7InByb2R1Y3QiOnsic2t1IjoiYXNkZiIsIm5hbWUiOiJuYW1lIiwicHJpY2UiOjEwMi4zLCJvcHRpb25zIjpbeyJrZXkiOiJzaXplIiwidmFsdWUiOiIxNSJ9XX0sIlF1YW50aXR5IjoxfV19" {
		t.Fatal("cookie error")
	}

	location, ok := rr.HeaderMap["Location"]
	if !ok {
		t.Fatal("location not found in the header")
	}
	locationString := location[0]
	if locationString != cart.Path {
		t.Fatal("location mismatch")
	}
}

func TestHandler_Delete_Item_Invalid(t *testing.T) {
	store := mem.NewCartStore("mcart")
	handler := cart.New(store)

	form := url.Values{}
	form.Add("_method", "delete")

	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("Received invalid code response: %d\n", rr.Code)
	}

	location, ok := rr.HeaderMap["Location"]
	if !ok {
		t.Fatal("location not found in the header")
	}
	locationString := location[0]
	if locationString != cart.Path {
		t.Fatal("location mismatch")
	}

	if rr.HeaderMap.Get("Set-Cookie") != "" {
		t.Fatal("cookie set?")
	}
}

func TestHandler_Delete_Item(t *testing.T) {
	store := mem.NewCartStore("mcart")
	handler := cart.New(store)

	sku := "xyz"
	price := 123.2
	i := model.CartItem{
		Product: model.Product{
			SKU:     sku,
			Name:    sku,
			Price:   price,
			Options: []model.ProductOption{},
		},
		Quantity: 1,
	}
	store.Cart.Items = append(store.Cart.Items, i)

	form := url.Values{}
	form.Add("_method", "delete")
	form.Add("sku", sku)

	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusFound {
		t.Fatalf("Received invalid code response: %d\n", rr.Code)
	}

	location, ok := rr.HeaderMap["Location"]
	if !ok {
		t.Fatal("location not found in the header")
	}
	locationString := location[0]
	if locationString != cart.Path {
		t.Fatal("location mismatch")
	}

	sh := strings.TrimPrefix(rr.HeaderMap.Get("Set-Cookie"), "mcart=")
	data, err := base64.StdEncoding.DecodeString(sh)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"shipping":0,"items":[]}` {
		t.Fatal("cart cookie error")
	}
}
