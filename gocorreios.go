package gocorreios

import (
	"github.com/Lgdev07/gocorreios/cep"
	"github.com/Lgdev07/gocorreios/fare"
	"github.com/Lgdev07/gocorreios/tracking"
)

// Tracking returns a well-formatted JSON response with tracking information.
func Tracking(codes []string) ([]byte, error) {
	body, err := tracking.TrackingResult(codes)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Cep returns a well-formatted JSON response with CEP information.
func Cep(cepCode string) ([]byte, error) {
	body, err := cep.CepResult(cepCode)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Fare returns a well-formatted JSON response with fare calculation information.
func Fare(fareInterf fare.Interface) ([]byte, error) {
	body, err := fare.FareResult(fareInterf)
	if err != nil {
		return nil, err
	}

	return body, nil
}
