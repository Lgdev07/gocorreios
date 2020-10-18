package tracking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

const (
	apiURL string = "http://webservice.correios.com.br/service/rest/rastro/rastroMobile"
	data   string = "<rastroObjeto><usuario></usuario><senha></senha><tipo>L</tipo><resultado>T</resultado><objetos>%v</objetos><lingua>101</lingua><token></token></rastroObjeto>"
)

// Object represents the main structure
type Object struct {
	Number      string    `json:"number"`
	Category    string    `json:"category"`
	Date        string    `json:"last_date"`
	Status      string    `json:"last_status"`
	Detail      string    `json:"last_detail"`
	Origin      string    `json:"last_origin"`
	Destination string    `json:"last_destination"`
	History     []History `json:"history"`
}

// History represents the array of events related to a Object
type History struct {
	Date        string `json:"date"`
	Status      string `json:"status"`
	Detail      string `json:"detail"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

// Tracking returns a well formatted json response of a tracking object
func Tracking(codes []string) ([]byte, error) {
	code := strings.Join(codes, "")

	body, err := searchCode(code)
	if err != nil {
		var b []byte
		return b, err
	}

	values := []Object{}

	objects := gjson.Get(string(body), "objeto")

	for _, objectValue := range objects.Array() {
		objeto := Object{}
		objeto.Number = gjson.Get(objectValue.String(), "numero").String()
		objeto.Category = gjson.Get(objectValue.String(), "categoria").String()

		events := gjson.Get(objectValue.String(), "evento")
		for key, event := range events.Array() {
			if key == 0 {

				objeto.Date = searchDate(event)
				objeto.Status = gjson.Get(event.String(), "descricao").String()
				objeto.Detail = gjson.Get(event.String(), "detalhe").String()
				objeto.Origin = searchOrigin(event)

				destinations := gjson.Get(event.String(), "destino")
				for _, destination := range destinations.Array() {
					objeto.Destination = searchDestination(destination)
				}

			} else {
				history := History{}

				history.Date = searchDate(event)
				history.Status = gjson.Get(event.String(), "descricao").String()
				history.Detail = gjson.Get(event.String(), "detalhe").String()
				history.Origin = searchOrigin(event)

				destinations := gjson.Get(event.String(), "destino")
				for _, destination := range destinations.Array() {
					history.Destination = searchDestination(destination)
				}

				objeto.History = append(objeto.History, history)
			}

		}
		values = append(values, objeto)
	}

	e, err := json.MarshalIndent(values, "", "    ")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return e, nil

}

// searchDestination returns the address of a destination
func searchDestination(destination gjson.Result) string {
	local := gjson.Get(destination.String(), "local")
	place := gjson.Get(destination.String(), "endereco.logradouro")
	number := gjson.Get(destination.String(), "endereco.numero")
	locality := gjson.Get(destination.String(), "endereco.localidade")
	uf := gjson.Get(destination.String(), "endereco.uf")
	neighborhood := gjson.Get(destination.String(), "endereco.bairro")

	return fmt.Sprintf("%v - %v, %v, %v - %v/%v", local, place, number, neighborhood, locality, uf)
}

// searchOrigin returns the address of an Origin
func searchOrigin(event gjson.Result) string {
	local := gjson.Get(event.String(), "unidade.local")
	place := gjson.Get(event.String(), "unidade.endereco.logradouro")
	number := gjson.Get(event.String(), "unidade.endereco.numero")
	locality := gjson.Get(event.String(), "unidade.endereco.localidade")
	uf := gjson.Get(event.String(), "unidade.endereco.uf")
	neighborhood := gjson.Get(event.String(), "unidade.endereco.bairro")

	return fmt.Sprintf("%v - %v, %v, %v - %v/%v", local, place, number, neighborhood, locality, uf)
}

// searchDate returns the date of an event
func searchDate(event gjson.Result) string {
	date := gjson.Get(event.String(), "data")
	hour := gjson.Get(event.String(), "hora")
	return fmt.Sprintf("%v - %v", date, hour)
}

// searchCode returns the body of a correios api requisition
func searchCode(_code string) ([]byte, error) {
	data := fmt.Sprintf(data, _code)
	client := &http.Client{}
	var b []byte

	r, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data))
	if err != nil {
		return b, err
	}

	r.Header.Add("Content-Type", "application/xml")

	resp, err := client.Do(r)
	if err != nil {
		return b, err
	}

	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return b, err
	}

	return _body, nil
}
