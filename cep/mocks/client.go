package mocks

import (
	"net/http"

	"github.com/Lgdev07/gocorreios/cep/client"
)

// MockGetType is the get type of ClientMock struct
type MockGetType func(url string) (*http.Response, error)

// ClientMock is the struct for our Mock
type ClientMock struct {
	MockGet MockGetType
}

var (
	// GetFunc is the variable to we make our mock
	GetFunc func(url string) (*http.Response, error)
)

// Get is the method oof ClientMock
func (c *ClientMock) Get(url string) (resp *http.Response, err error) {
	return GetFunc(url)
}

func init() {
	client.Client = &ClientMock{}
}
