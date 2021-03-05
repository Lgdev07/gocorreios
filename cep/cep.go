package cep

import (
	"encoding/json"

	"github.com/Lgdev07/gocorreios/cep/services"
)

// CepResult returns a well formatted json response of a cep object
func CepResult(cep string) ([]byte, error) {
	var b []byte

	result := getFirstResult(cep)
	if result.Err != nil {
		return b, result.Err
	}

	identedCepItem, err := json.MarshalIndent(result.Res, "", "    ")
	if err != nil {
		return b, err
	}

	return identedCepItem, nil
}

func getFirstResult(cep string) services.ResultError {
	resultChan := make(chan services.ResultError)
	quit := make(chan bool)
	var err error

	result := services.ResultError{
		Res: services.Item{},
		Err: err,
	}

	go runWorkers(cep, resultChan, quit)

	for {
		select {
		case v := <-resultChan:
			if v.Err != nil {
				result.Err = v.Err
				continue
			}
			return v
		}
	}
}

func runWorkers(cep string, resultChan chan services.ResultError, quit chan bool) {
	go func() { services.SearchCepViaCEPAPI(cep, resultChan) }()
	go func() { services.SearchBrasilApi(cep, resultChan) }()
}
