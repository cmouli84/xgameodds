package infrastructure

import (
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
}

// GetHTTPResponse function
func GetHTTPResponse(url string) []byte {
	resp, err := client.Get(url)
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
