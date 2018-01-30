package main

import (
	"github.com/go-chi/chi"
	"github.com/miaversa/backend/router"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var VERSION string
var BRANCH string
var COMMIT string

var log = logrus.New()

func main() {
	log.Formatter = &logrus.TextFormatter{}
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)

	contextLogger := log.WithFields(logrus.Fields{"environment": "dev"})
	contextLogger.Info("init")

	r := chi.NewRouter()
	r.Mount("/carrinho", router.CartRouter())
	r.Mount("/cliente", router.CustomerRouter())
	r.Mount("/pagamento", router.PaymentRouter())

	log.Error(http.ListenAndServe(":8080", r))
}
