package http_helper

import (
	"io/ioutil"
	"log"
	"net/http"
)

type RequestApi interface {
	RequestGetApi(url string) ([]byte, error)
}
type RequestApiClient struct {
}

func NewRequestApiClient() *RequestApiClient {
	return &RequestApiClient{}
}

func (RequestApiClient) RequestGetApi(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return responseData, nil
}
func InternetConnected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}