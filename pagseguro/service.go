package pagseguro

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/miaversa/backend/config"
)

var SID_REGEXP = regexp.MustCompile("<id>(.*?)</id>")

type PagSeguroService interface {
	SessionID() (string, error)
}

type pagSeguro struct{}

func New() *pagSeguro {
	return &pagSeguro{}
}

func (s *pagSeguro) SessionID() (string, error) {
	baseURL := config.PagSeguroBaseAPI + "sessions/"

	log.Println(baseURL)
	log.Println(config.PagSeguroEmail)
	log.Println(config.PagSeguroToken)

	form := url.Values{}
	form.Add("email", config.PagSeguroEmail)
	form.Add("token", config.PagSeguroToken)

	resp, err := http.PostForm(baseURL, form)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("status not 200")
		log.Println(resp.StatusCode)
		return "", errors.New("status code error")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return "", err
	}

	bodyString := string(bodyBytes)

	rslice := SID_REGEXP.FindStringSubmatch(bodyString)
	if len(rslice) < 2 {
		return "", errors.New("faltou session id")
	}

	return rslice[1], nil
}
