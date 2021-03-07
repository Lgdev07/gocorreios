package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/Lgdev07/gocorreios/http/client"
)

// SearchBrasilApi makes a request to brasil api and return the body of the response
func SearchBrasilApi(cepString string, resultChan chan ResultError) {
	cepAPIURL := "https://brasilapi.com.br/api/cep/v1/%v"
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

	var dat map[string]interface{}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		resultChan <- ResultError{Res: *cepItem, Err: err}
		return
	}

	if dat["errors"] != nil {
		firstError := getFirstErrorMessage(dat["errors"])
		resultChan <- ResultError{Res: *cepItem, Err: firstError}
		return
	}

	cepItem = &Item{
		Cep:          dat["cep"].(string),
		City:         dat["city"].(string),
		State:        dat["state"].(string),
		Street:       dat["street"].(string),
		Neighborhood: dat["neighborhood"].(string),
	}

	resultChan <- ResultError{Res: *cepItem, Err: nil}
	return
}

func getFirstErrorMessage(errorInterface interface{}) error {
	errorSlice := errorInterface.([]interface{})
	newError := errors.New(errorSlice[0].(map[string]interface{})["message"].(string))
	return newError
}
