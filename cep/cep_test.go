package cep

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Lgdev07/gocorreios/cep/mocks"
	"github.com/tj/assert"
)

const (
	cepNumber string = "01311-915"
)

func TestCepResult(t *testing.T) {
	jsonResponse := `{
		"cep": "099811380",
		"uf": "SP",
		"localidade": "São Paulo",
		"logradouro": "Rua Maria de Lourdes",
		"bairro": "Jardim Ruyce",
		"complemento": "",
		"city": "São Paulo",
		"state": "SP",
		"street": "Rua Maria de Lourdes",
		"neighborhood": "Jardim Ruyce"
	 }`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	mocks.GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	var dat map[string]interface{}

	result, _ := CepResult("099811380")

	_ = json.Unmarshal(result, &dat)

	assert.Equal(t, len(dat), 6)
	assert.Equal(t, dat["cep"], "099811380")
	assert.Equal(t, dat["uf"], "SP")
	assert.Equal(t, dat["localidade"], "São Paulo")
	assert.Equal(t, dat["logradouro"], "Rua Maria de Lourdes")
	assert.Equal(t, dat["bairro"], "Jardim Ruyce")
	assert.Equal(t, dat["complemento"], "")
}

func TestCepNotExist(t *testing.T) {
	jsonResponse := `{
		<>,
		"errors": [{
			"message": "CEP informado possui mais do que 8 caracteres."
		}]
	 }`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	mocks.GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 500,
			Body:       r,
		}, nil
	}

	_, err := CepResult("Wrong number")
	assert.Error(t, err)
}
