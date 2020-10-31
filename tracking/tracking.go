package tracking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

const (
	trackingAPIURL string = "http://webservice.correios.com.br/service/rest/rastro/rastroMobile"
	trackingData   string = "<rastroObjeto><usuario></usuario><senha></senha><tipo>L</tipo><resultado>T</resultado><objetos>%v</objetos><lingua>101</lingua><token></token></rastroObjeto>"
)

// Object represents the main structure
type Object struct {
	Number      string    `json:"number"`
	Category    string    `json:"category"`
	Date        string    `json:"last_date"`
	Type        string    `json:"last_type"`
	Status      int       `json:"last_status"`
	Description string    `json:"last_description"`
	Detail      string    `json:"last_detail"`
	Origin      string    `json:"last_origin"`
	Destination string    `json:"last_destination"`
	History     []History `json:"history"`
}

// History represents the array of events related to a Object
type History struct {
	Date        string `json:"date"`
	Type        string `json:"type"`
	Status      int    `json:"status"`
	Description string `json:"description"`
	Detail      string `json:"detail"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

// TrackingResult returns a well formatted json response of a tracking object
func TrackingResult(codes []string) ([]byte, error) {
	code := strings.Join(codes, "")

	body, err := searchCode(code)
	if err != nil {
		var b []byte
		return b, err
	}

	objects := getObjects(body)

	e, err := json.MarshalIndent(objects, "", "    ")
	if err != nil {
		var b []byte
		return b, err
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
	var b []byte

	trackingData := fmt.Sprintf(trackingData, _code)
	client := &http.Client{}

	r, err := http.NewRequest("POST", trackingAPIURL, bytes.NewBufferString(trackingData))
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

// getObjects returns the struct with after parse from json
func getObjects(body []byte) []Object {
	var values []Object

	objects := gjson.Get(string(body), "objeto")

	for _, objectValue := range objects.Array() {
		var object Object
		object.Number = gjson.Get(objectValue.String(), "numero").String()
		object.Category = gjson.Get(objectValue.String(), "categoria").String()

		events := gjson.Get(objectValue.String(), "evento")
		for key, event := range events.Array() {
			if key == 0 {

				object.Date = searchDate(event)
				object.Type = gjson.Get(event.String(), "tipo").String()
				object.Status = int(gjson.Get(event.String(), "status").Int())
				object.Description = gjson.Get(event.String(), "descricao").String()
				object.Detail = gjson.Get(event.String(), "detalhe").String()
				object.Origin = searchOrigin(event)

				destinations := gjson.Get(event.String(), "destino")
				for _, destination := range destinations.Array() {
					object.Destination = searchDestination(destination)
				}

			} else {
				history := History{}

				history.Date = searchDate(event)
				history.Type = gjson.Get(event.String(), "tipo").String()
				history.Status = int(gjson.Get(event.String(), "status").Int())
				history.Description = gjson.Get(event.String(), "descricao").String()
				history.Detail = gjson.Get(event.String(), "detalhe").String()
				history.Origin = searchOrigin(event)

				destinations := gjson.Get(event.String(), "destino")
				for _, destination := range destinations.Array() {
					history.Destination = searchDestination(destination)
				}

				object.History = append(object.History, history)
			}

		}
		values = append(values, object)
	}

	return values
}
