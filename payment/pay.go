package payment

import (
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
)

var Path string = "/pagar"
var templateFile string = "payment.html"

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.New(templateFile)
	t.Parse(string(templates.MustAsset(templateFile)))
	t.Execute(w, nil)
}
