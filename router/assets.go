package router

import (
	"github.com/go-chi/chi"
	"github.com/miaversa/backend/assets"
	"net/http"
	"strings"
)

func AssetsRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/{filename}", sendAsset)
	return r
}

func sendAsset(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")

	if strings.HasSuffix(filename, ".css") {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}

	w.Write([]byte(assets.MustAsset(filename)))
}
