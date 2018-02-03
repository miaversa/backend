package assets

import (
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

var Path string = "/assets/{filename}"

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	if strings.HasSuffix(filename, ".css") {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}
	w.Write([]byte(MustAsset(filename)))
}
