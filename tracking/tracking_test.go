package tracking

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Lgdev07/gocorreios/http/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	codeNumber          string = "ON732904576BR"
	codeNotFoundMessage string = "ERRO: Objeto não encontrado na base de dados dos Correios."
)

func TestSearchCodeExists(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResult)))

	mocks.GetDo = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := getBody(codeNumber)
	if err != nil {
		t.Errorf("Error while getting getBody: %v", err)
	}
	assert.Contains(t, string(response), codeNumber)
}

func TestSearchCodeNotExist(t *testing.T) {
	jsonResult := `{
    "objeto": [
        {
            "categoria": "ERRO: Objeto não encontrado na base de dados dos Correios.",
            "numero": "ON732904576B"
        }
    ],
    "pesquisa": "Lista de Objetos",
    "quantidade": "1",
    "resultado": "Todos os eventos",
    "versao": "3.0"
	}`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResult)))

	mocks.GetDo = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := getBody("WrongCode")
	if err != nil {
		t.Errorf("Error while getting getBody: %v", err)
	}
	assert.Contains(t, string(response), codeNotFoundMessage)
}

func TestTracking(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResult)))

	mocks.GetDo = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := TrackingResult([]string{codeNumber})
	if err != nil {
		t.Errorf("Error while getting Tracking: %v", err)
	}

	var responseInterface []map[string]interface{}
	err = json.Unmarshal(response, &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, len(responseInterface), 1)

	assert.Contains(t, responseInterface[0], "category")
	assert.Contains(t, responseInterface[0], "last_date")
	assert.Contains(t, responseInterface[0], "last_type")
	assert.Contains(t, responseInterface[0], "last_status")
	assert.Contains(t, responseInterface[0], "last_description")
	assert.Contains(t, responseInterface[0], "last_detail")
	assert.Contains(t, responseInterface[0], "last_origin")
	assert.Contains(t, responseInterface[0], "last_destination")
	assert.Contains(t, responseInterface[0], "history")

	assert.Equal(t, responseInterface[0]["number"], codeNumber)
	assert.Equal(t, responseInterface[0]["last_date"], "05/03/2021 - 14:04")
	assert.Equal(t, responseInterface[0]["last_type"], "BDE")
	assert.Equal(t, responseInterface[0]["last_description"], "Objeto entregue ao destinatário")
	assert.Equal(t, responseInterface[0]["last_detail"], "")
	assert.Equal(t, responseInterface[0]["last_origin"], "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP")
	assert.Equal(t, responseInterface[0]["last_destination"], "")
}

func TestTrackingNotExist(t *testing.T) {
	jsonResult := `{
    "objeto": [
        {
            "categoria": "ERRO: Objeto não encontrado na base de dados dos Correios.",
            "numero": "ON732904576B"
        }
    ],
    "pesquisa": "Lista de Objetos",
    "quantidade": "1",
    "resultado": "Todos os eventos",
    "versao": "3.0"
	}`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResult)))

	mocks.GetDo = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := TrackingResult([]string{"Wrong number"})
	if err != nil {
		t.Errorf("Error while getting Tracking: %v", err)
	}

	var responseInterface []map[string]interface{}
	err = json.Unmarshal(response, &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, len(responseInterface), 1)
	assert.Equal(t, responseInterface[0]["error"], codeNotFoundMessage)
}
