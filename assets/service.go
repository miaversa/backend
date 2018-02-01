package assets

import (
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	if strings.HasSuffix(filename, ".css") {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}
	w.Write([]byte(MustAsset(filename)))
}
