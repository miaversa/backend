package main

import (
	"github.com/go-chi/chi"
	"github.com/miaversa/backend/router"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var log = logrus.New()

func main() {
	log.Formatter = &logrus.TextFormatter{}
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)

	contextLogger := log.WithFields(logrus.Fields{"environment": "dev"})
	contextLogger.Info("init")

	r := chi.NewRouter()
	r.Mount("/carrinho", router.CartRouter())

	log.Error(http.ListenAndServe(":8080", r))
}
