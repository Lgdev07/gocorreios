package mocks

import (
	"net/http"

	"github.com/Lgdev07/gocorreios/http/client"
)

// ClientMock is the struct for our Mock
type ClientMock struct{}

var (
	// GetFunc is the variable to we make our mock
	GetFunc func(url string) (*http.Response, error)

	// GetDo is the variable to we make our mock
	GetDo func(req *http.Request) (*http.Response, error)
)

// Get is the method of ClientMock
func (c *ClientMock) Get(url string) (resp *http.Response, err error) {
	return GetFunc(url)
}

// Do is the method of ClientMock
func (c *ClientMock) Do(req *http.Request) (resp *http.Response, err error) {
	return GetDo(req)
}

func init() {
	client.Client = &ClientMock{}
}
