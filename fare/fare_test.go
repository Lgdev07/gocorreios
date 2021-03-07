package fare

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Lgdev07/gocorreios/http/mocks"
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
	xmlResponse := `<?xml version="1.0" encoding="utf-8"?>
	<cResultado xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://tempuri.org/">
		<Servicos>
			<cServico>
				<Codigo>40010</Codigo>
				<Valor>843,50</Valor>
				<PrazoEntrega>3</PrazoEntrega>
				<ValorMaoPropria>0,00</ValorMaoPropria>
				<ValorAvisoRecebimento>0,00</ValorAvisoRecebimento>
				<ValorValorDeclarado>0,00</ValorValorDeclarado>
				<EntregaDomiciliar>S</EntregaDomiciliar>
				<EntregaSabado>S</EntregaSabado>
				<Erro>0</Erro>
				<MsgErro />
				<ValorSemAdicionais>843,50</ValorSemAdicionais>
				<obsFim />
			</cServico>
		</Servicos>
	</cResultado>`

	r := ioutil.NopCloser(bytes.NewReader([]byte(xmlResponse)))

	mocks.GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

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

	assert.Equal(t, "SEDEX", responseInterface["service"])
	assert.Equal(t, "843,50", responseInterface["price"])
	assert.Equal(t, "3", responseInterface["days_for_delivery"])
	assert.Equal(t, "S", responseInterface["deliver_home"])
	assert.Equal(t, "S", responseInterface["deliver_saturday"])
}

func TestFareResultWrongService(t *testing.T) {
	params.Service = "Wrong Service"
	_, err := FareResult(params)

	assert.Error(t, err)
	assert.Equal(t, "Avaliable Services: PAC, SEDEX", err.Error())
}

func TestFareResultWrongCepOrigin(t *testing.T) {
	xmlResponse := `<?xml version="1.0" encoding="utf-8"?>
		<cResultado xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://tempuri.org/">
			<Servicos>
				<cServico>
					<Codigo>40010</Codigo>
					<Valor>0,00</Valor>
					<PrazoEntrega>0</PrazoEntrega>
					<ValorMaoPropria>0,00</ValorMaoPropria>
					<ValorAvisoRecebimento>0,00</ValorAvisoRecebimento>
					<ValorValorDeclarado>0,00</ValorValorDeclarado>
					<EntregaDomiciliar />
					<EntregaSabado />
					<Erro>-2</Erro>
					<MsgErro>CEP de origem invalido.</MsgErro>
					<ValorSemAdicionais>0,00</ValorSemAdicionais>
					<obsFim />
				</cServico>
			</Servicos>
		</cResultado>`

	r := ioutil.NopCloser(bytes.NewReader([]byte(xmlResponse)))

	mocks.GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	params.Service = "SEDEX"

	_, err := FareResult(params)
	assert.Error(t, err)
	assert.Equal(t, "CEP de origem invalido.", err.Error())
}

func TestFareResultWrongCepDestination(t *testing.T) {
	xmlResponse := `<?xml version="1.0" encoding="utf-8"?>
		<cResultado xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://tempuri.org/">
			<Servicos>
				<cServico>
					<Codigo>40010</Codigo>
					<Valor>0,00</Valor>
					<PrazoEntrega>0</PrazoEntrega>
					<ValorMaoPropria>0,00</ValorMaoPropria>
					<ValorAvisoRecebimento>0,00</ValorAvisoRecebimento>
					<ValorValorDeclarado>0,00</ValorValorDeclarado>
					<EntregaDomiciliar />
					<EntregaSabado />
					<Erro>-2</Erro>
					<MsgErro>CEP de destino invalido.</MsgErro>
					<ValorSemAdicionais>0,00</ValorSemAdicionais>
					<obsFim />
				</cServico>
			</Servicos>
		</cResultado>`

	r := ioutil.NopCloser(bytes.NewReader([]byte(xmlResponse)))

	mocks.GetFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	params.Service = "SEDEX"

	_, err := FareResult(params)
	assert.Error(t, err)
	assert.Equal(t, "CEP de destino invalido.", err.Error())
}
