package infrastructure

import (
	"io/ioutil"
	"net/http"
	"time"
)

// NewHTTPClient returns new Http client
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
}

// HTTPClientInterface interface
type HTTPClientInterface interface {
	GetHTTPResponse(url string) []byte
}

// HTTPClientHandler struct
type HTTPClientHandler struct {
	Client http.Client
}

// GetHTTPResponse function
func (httpClientHandler *HTTPClientHandler) GetHTTPResponse(url string) []byte {
	resp, err := httpClientHandler.Client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, ioerr := ioutil.ReadAll(resp.Body)
	if ioerr != nil {
		panic(ioerr)
	}
	return response
}
