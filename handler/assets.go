package handler

import (
	"github.com/miaversa/backend/assets"
	"net/http"
	"strings"
)

// ServeHTTP sends the file with the correct content type.
func AssetHandler(w http.ResponseWriter, r *http.Request) error {
	filename := r.URL.Query().Get("filename")
	if strings.HasSuffix(filename, ".css") {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}
	a, err := assets.Asset(filename)
	if err != nil {
		return err
	}
	w.Write(a)
	return nil
}
