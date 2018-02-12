package handler

import (
	"github.com/miaversa/backend/session"
	"net/http"
)

type redirectHandler struct {
	sessionStorage session.SessionService
}

func NewRedirectHandler(sessionStorage session.SessionService) *redirectHandler {
	return &redirectHandler{sessionStorage}
}

func (rh *redirectHandler) RedirectIfSession(h HandlerFuncErr) HandlerFuncErr {
	return func(w http.ResponseWriter, r *http.Request) error {
		r.ParseForm()
		redirect := ""
		if r.FormValue("redirect") != "" {
			redirect = r.FormValue("redirect")
		}
		session, err := rh.sessionStorage.Get()
		if err != nil {
			return err
		}
		if session != "" {
			http.Redirect(w, r, redirect, http.StatusFound)
			return nil
		}
		return nil
	}
}
