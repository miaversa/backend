package handler

import (
	"net/http"
)

func HandlerError(h func(w http.ResponseWriter, r *http.Request) error) {
}
