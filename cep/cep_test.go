package cep

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	cepNumber          string = "01311915"
	cepNotFoundMessage string = "Invalid CEP"
)

func TestSearchCepExists(t *testing.T) {
	response, err := searchCepBrasilAPI(cepNumber)
	if err != nil {
		t.Errorf("Error while getting searchCepBrasilAPI: %v", err)
	}
	assert.Equal(t, response.Cep, cepNumber)
}

func TestSearchCepNotExist(t *testing.T) {
	_, err := searchCepBrasilAPI("WrongCode")
	assert.Error(t, err)
}

func TestCepResult(t *testing.T) {
	response, err := CepResult(cepNumber)
	if err != nil {
		t.Errorf("Error while getting Tracking: %v", err)
	}

	var responseInterface map[string]interface{}
	err = json.Unmarshal(response, &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, responseInterface["cep"], cepNumber)
	assert.Equal(t, responseInterface["state"], "SP")
	assert.Equal(t, responseInterface["city"], "São Paulo")
	assert.Equal(t, responseInterface["street"], "Avenida Paulista 807")
	assert.Equal(t, responseInterface["neighborhood"], "Bela Vista")
}

func TestTrackingNotExist(t *testing.T) {
	_, err := CepResult("Wrong number")
	assert.Error(t, err)

}
