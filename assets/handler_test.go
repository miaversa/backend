package assets_test

import (
	"github.com/miaversa/backend/assets"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Valid(t *testing.T) {
	handler := assets.New()
	req, _ := http.NewRequest(http.MethodGet, "/assets?filename=tachyons.min.css", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("Received non-200 response: %d\n", rr.Code)
	}

	req, _ = http.NewRequest(http.MethodGet, "/assets?filename=x.min.css", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Fatalf("Received non-404 response: %d\n", rr.Code)
	}
}
