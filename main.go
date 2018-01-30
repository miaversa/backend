package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/miaversa/backend/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var VERSION string
var BRANCH string
var COMMIT string

var log = logrus.New()

func main() {
	viper.SetDefault("license", "apache")

	log.Formatter = &logrus.TextFormatter{}
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)

	contextLogger := log.WithFields(logrus.Fields{"environment": "dev"})
	contextLogger.Info("init")

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	r.Mount("/carrinho", router.CartRouter())
	r.Mount("/cliente", router.CustomerRouter())
	r.Mount("/pagamento", router.PaymentRouter())
	r.Mount("/assets", router.AssetsRouter())

	log.Error(http.ListenAndServe(":8080", r))
}
