package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/Lgdev07/gocorreios/http/client"
)

// SearchCepViaCEPAPI makes a request to via cep api and return the body of the response
func SearchCepViaCEPAPI(cepString string, resultChan chan ResultError) {
	cepAPIURL := "https://viacep.com.br/ws/%v/json/"
	cepItem := &Item{}

	url := fmt.Sprintf(cepAPIURL, cepString)

	resp, err := client.Client.Get(url)
	if err != nil {
		resultChan <- ResultError{Res: *cepItem, Err: err}
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resultChan <- ResultError{Res: *cepItem, Err: err}
		return
	}

	err = json.Unmarshal(body, cepItem)
	if err != nil {
		if err.Error() == "invalid character '<' looking for beginning of value" {
			resultChan <- ResultError{Res: *cepItem, Err: errors.New(
				"Please inform CEP in right format: 00000-000 or 00000000")}
			return
		}
		resultChan <- ResultError{Res: *cepItem, Err: err}
		return
	}

	if cepItem.Cep == "" {
		resultChan <- ResultError{Res: *cepItem, Err: errors.New("Invalid CEP")}
		return
	}

	resultChan <- ResultError{Res: *cepItem, Err: nil}
	return
}
