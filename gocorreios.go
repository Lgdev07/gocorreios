package gocorreios

import (
	"github.com/Lgdev07/gocorreios/cep"
	"github.com/Lgdev07/gocorreios/fare"
	"github.com/Lgdev07/gocorreios/tracking"
)

//Tracking returns a well formatted json response with tracking information
func Tracking(codes []string) ([]byte, error) {
	body, err := tracking.TrackingResult(codes)
	if err != nil {
		var b []byte
		return b, err
	}

	return body, nil
}

//Cep returns a well formatted json response with cep information
func Cep(cepCode string) ([]byte, error) {
	body, err := cep.CepResult(cepCode)
	if err != nil {
		var b []byte
		return b, err
	}

	return body, nil
}

//Fare returns a well formatted json response with fare calculatio information
func Fare(fareInterf fare.Interface) ([]byte, error) {
	body, err := fare.FareResult(fareInterf)
	if err != nil {
		var b []byte
		return b, err
	}

	return body, nil
}
