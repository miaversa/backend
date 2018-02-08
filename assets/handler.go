package assets

import (
	"net/http"
	"strings"
)

// Path for the routing
var Path string = "/assets/{filename}"

type handler struct {
}

// New creates a new assets handler
func New() *handler {
	return &handler{}
}

// ServeHTTP sends the file with the correct content type.
func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if strings.HasSuffix(r.URL.Path, ".css") {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}
	a, err := Asset(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Write(a)
}
