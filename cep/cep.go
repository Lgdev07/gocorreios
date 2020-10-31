package cep

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Item represents the cep structure
type Item struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
}

// CepResult returns a well formatted json response of a cep object
func CepResult(cep string) ([]byte, error) {
	var b []byte

	item, err := searchCepBrasilAPI(cep)
	if err != nil {
		return b, err
	}

	e, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		return b, err
	}

	return e, nil
}

// SearchCepBrasilAPI makes a request to brasil api and return the body of the response
func searchCepBrasilAPI(cepString string) (*Item, error) {
	cepAPIURL := "https://brasilapi.com.br/api/cep/v1/%v"
	cepItem := &Item{}

	url := fmt.Sprintf(cepAPIURL, cepString)

	resp, err := http.Get(url)
	if err != nil {
		return cepItem, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cepItem, err
	}

	err = json.Unmarshal(body, cepItem)
	if err != nil {
		return cepItem, err
	}

	if cepItem.Cep == "" {
		return cepItem, errors.New("Invalid CEP")
	}

	return cepItem, nil
}
