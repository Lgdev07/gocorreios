package fare

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
)

const (
	fareURL      string = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx/CalcPrecoPrazo?"
	fareQueryURL string = `nCdEmpresa=&sDsSenha=&nCdServico=%v&sCepOrigem=%v&sCepDestino=%v&nVlPeso=%v&nCdFormato=1&nVlComprimento=%v&nVlAltura=%v&nVlLargura=%v&sCdMaoPropria=n&nVlValorDeclarado=0&sCdAvisoRecebimento=n&nVlDiametro=0&StrRetorno=xml`
)

var avaliableServices map[string]string = map[string]string{"PAC": "41106", "SEDEX": "40010"}

//Interface represents the interface of fare request arguments
type Interface struct {
	Service        string
	CepOrigin      string
	CepDestination string
	Weight         float64
	Length         float64
	Height         float64
	Width          float64
}

//Item represents the struct that we receive from correios service
type item struct {
	Servico           string `json:"service"`
	Valor             string `xml:"Servicos>cServico>Valor" json:"price"`
	Prazo             string `xml:"Servicos>cServico>PrazoEntrega" json:"days_for_delivery"`
	EntregaDomiciliar string `xml:"Servicos>cServico>EntregaDomiciliar" json:"deliver_home"`
	EntregaSabado     string `xml:"Servicos>cServico>EntregaSabado" json:"deliver_saturday"`
	Erro              string `xml:"Servicos>cServico>Erro" json:"-"`
	MsgErro           string `xml:"Servicos>cServico>MsgErro" json:"obs"`
}

// FareResult returns a well formatted json response of a fare object
func FareResult(query Interface) ([]byte, error) {
	var b []byte
	service := avaliableServices[query.Service]
	if service == "" {
		return b, errors.New("Avaliable Services: PAC, SEDEX")
	}

	queryValues := fmt.Sprintf(fareQueryURL, service, query.CepOrigin, query.CepDestination, query.Weight, query.Length, query.Height, query.Width)
	resp, err := http.Get(fareURL + queryValues)
	if err != nil {
		return b, err
	}
	defer resp.Body.Close()

	var item item
	decode := xml.NewDecoder(resp.Body)
	err = decode.Decode(&item)
	if err != nil {
		return b, err
	}

	if item.Valor == "0,00" {
		return b, errors.New(item.MsgErro)
	}

	item.Servico = query.Service
	e, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		return b, err
	}

	return e, nil
}
