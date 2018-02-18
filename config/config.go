package config

import (
	"os"
)

var Port string
var Sandbox = true
var PagSeguroBaseAPI = "https://ws.sandbox.pagseguro.uol.com.br/v2/"
var PagSeguroJavascript = "https://stc.sandbox.pagseguro.uol.com.br/pagseguro/api/v2/checkout/pagseguro.directpayment.js"

var PagSeguroEmail string
var PagSeguroToken string

func Load() {
	Port = os.Getenv("BKND_PORT")
	PagSeguroEmail = os.Getenv("BKND_PGS_EMAIL")
	PagSeguroToken = os.Getenv("BKND_PGS_TOKEN")
}
