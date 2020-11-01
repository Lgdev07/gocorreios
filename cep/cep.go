package cep

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Item represents the cep structure
type item struct {
	Cep          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	Complement   string `json:"complemento"`
}

// CepResult returns a well formatted json response of a cep object
func CepResult(cep string) ([]byte, error) {
	var b []byte

	item, err := searchCepViaCEPAPI(cep)
	if err != nil {
		return b, err
	}

	e, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		return b, err
	}

	return e, nil
}

// searchCepViaCEPAPI makes a request to via cep api and return the body of the response
func searchCepViaCEPAPI(cepString string) (*item, error) {
	cepAPIURL := "https://viacep.com.br/ws/%v/json/"
	cepItem := &item{}

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
		if err.Error() == "invalid character '<' looking for beginning of value" {
			return cepItem, errors.New("Please inform CEP in right format: 00000-000 or 00000000")
		}
		return cepItem, err
	}

	if cepItem.Cep == "" {
		return cepItem, errors.New("Invalid CEP")
	}

	return cepItem, nil
}
