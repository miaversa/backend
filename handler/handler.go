package handler

import (
	"net/http"
)

type HandlerFuncErr func(http.ResponseWriter, *http.Request) error
