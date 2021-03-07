package cep

import (
	"encoding/json"

	"github.com/Lgdev07/gocorreios/cep/services"
)

var servicesList []string = []string{"brasilApi", "viaCepApi"}

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
	var errorCount int
	resultChan := make(chan services.ResultError)
	result := services.ResultError{
		Res: services.Item{},
	}

	go runServices(cep, resultChan)

	for {
		select {
		case v := <-resultChan:
			if v.Err != nil {
				if returnError(&errorCount) {
					result.Err = v.Err
					return result
				}
				continue
			}
			return v
		}
	}

}

func returnError(errorCount *int) bool {
	var lenServices int = len(servicesList)

	*errorCount++
	if *errorCount == lenServices {
		return true
	}

	return false
}

func runServices(cep string, resultChan chan services.ResultError) {
	go func() { services.SearchCepViaCEPAPI(cep, resultChan) }()
	go func() { services.SearchBrasilApi(cep, resultChan) }()
}
