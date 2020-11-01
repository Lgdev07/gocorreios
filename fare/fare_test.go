package fare

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var params = Interface{
	Service:        "SEDEX",
	CepOrigin:      "60010100",
	CepDestination: "04029000",
	Weight:         30,
	Length:         20,
	Height:         20,
	Width:          20,
}

func TestFareResultExists(t *testing.T) {
	response, err := FareResult(params)
	if err != nil {
		t.Errorf("Error while getting TestFareResult: %v", err)
	}

	var responseInterface map[string]interface{}
	err = json.Unmarshal(response, &responseInterface)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}

	assert.Contains(t, responseInterface, "service")
	assert.Contains(t, responseInterface, "price")
	assert.Contains(t, responseInterface, "days_for_delivery")
	assert.Contains(t, responseInterface, "deliver_home")
	assert.Contains(t, responseInterface, "deliver_saturday")

	assert.NotEmpty(t, responseInterface["service"])
	assert.NotEmpty(t, responseInterface["price"])
	assert.NotEmpty(t, responseInterface["days_for_delivery"])
	assert.NotEmpty(t, responseInterface["deliver_home"])
	assert.NotEmpty(t, responseInterface["deliver_saturday"])
}

func TestTestFareResultNotExists(t *testing.T) {
	params.Service = "Wrong Service"
	_, err := FareResult(params)
	assert.Error(t, err)
}
