package tracking

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	codeNumber          string = "OK816158697BR"
	codeNotFoundMessage string = "ERRO: Objeto não encontrado na base de dados dos Correios."
)

func TestSearchCodeExists(t *testing.T) {
	response, err := searchCode(codeNumber)
	if err != nil {
		t.Errorf("Error while getting searchCode: %v", err)
	}
	assert.Contains(t, string(response), codeNumber)
}

func TestSearchCodeNotExist(t *testing.T) {
	response, err := searchCode("WrongCode")
	if err != nil {
		t.Errorf("Error while getting searchCode: %v", err)
	}
	assert.Contains(t, string(response), codeNotFoundMessage)
}

func TestTracking(t *testing.T) {
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
	assert.Equal(t, responseInterface[0]["number"], codeNumber)
	assert.Contains(t, responseInterface[0], "category")
	assert.Contains(t, responseInterface[0], "last_date")
	assert.Contains(t, responseInterface[0], "last_type")
	assert.Contains(t, responseInterface[0], "last_status")
	assert.Contains(t, responseInterface[0], "last_description")
	assert.Contains(t, responseInterface[0], "last_detail")
	assert.Contains(t, responseInterface[0], "last_origin")
	assert.Contains(t, responseInterface[0], "last_destination")
	assert.Contains(t, responseInterface[0], "history")
}

func TestTrackingNotExist(t *testing.T) {
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
	assert.Equal(t, responseInterface[0]["category"], codeNotFoundMessage)
}

func TestTrackingMultipleObjects(t *testing.T) {
	response, err := TrackingResult([]string{codeNumber, codeNumber})
	if err != nil {
		t.Errorf("Error while getting Tracking: %v", err)
	}

	var responseInterface []map[string]interface{}
	err = json.Unmarshal(response, &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, len(responseInterface), 2)
}
