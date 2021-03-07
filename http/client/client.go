package client

import "net/http"

// HTTPClient is the interface for HTTPClient
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
	Do(req *http.Request) (*http.Response, error)
}

var (
	// Client is the struct for our HTTPClient
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}
