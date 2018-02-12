package handler_test

import (
	"github.com/miaversa/backend/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAssetHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/assets?filename=tachyons.min.css", nil)
	rr := httptest.NewRecorder()
	err := handler.AssetHandler(rr, req)
	if err != nil {
		t.Fatal(err)
	}
	if rr.Code != http.StatusOK {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}
	req, _ = http.NewRequest(http.MethodGet, "/assets?filename=x.min.css", nil)
	rr = httptest.NewRecorder()
	err = handler.AssetHandler(rr, req)
	if err == nil {
		t.Fatal("um erro era esperado")
	}
}
