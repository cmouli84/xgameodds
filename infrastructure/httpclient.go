package infrastructure

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// NewHTTPClient returns new Http client
func NewHTTPClient() *HTTPClientHandler {
	httpClientHandler := &HTTPClientHandler{}
	client := http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	httpClientHandler.Client = client

	return httpClientHandler
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
		fmt.Println(err)
		return []byte{}
	}
	defer resp.Body.Close()

	response, ioerr := ioutil.ReadAll(resp.Body)
	if ioerr != nil {
		fmt.Println(ioerr)
		return []byte{}
	}
	return response
}
