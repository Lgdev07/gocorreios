package tracking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Lgdev07/gocorreios/http/client"
	"github.com/tidwall/gjson"
)

const (
	trackingAPIURL string = "http://webservice.correios.com.br/service/rest/rastro/rastroMobile"
	trackingData   string = "<rastroObjeto><usuario></usuario><senha></senha><tipo>L</tipo><resultado>T</resultado><objetos>%v</objetos><lingua>101</lingua><token></token></rastroObjeto>"
)

type object struct {
	Number      string    `json:"number"`
	Category    string    `json:"category"`
	Date        string    `json:"last_date"`
	Type        string    `json:"last_type"`
	Status      int       `json:"last_status"`
	Description string    `json:"last_description"`
	Detail      string    `json:"last_detail"`
	Origin      string    `json:"last_origin"`
	Destination string    `json:"last_destination"`
	Error       string    `json:"error"`
	History     []history `json:"history"`
}

type history struct {
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

	body, err := getBody(code)
	if err != nil {
		var b []byte
		return b, err
	}

	objects := formatBody(body)

	e, err := json.MarshalIndent(objects, "", "    ")
	if err != nil {
		var b []byte
		return b, err
	}

	return e, nil
}

func getBody(_code string) ([]byte, error) {
	var b []byte

	trackingData := fmt.Sprintf(trackingData, _code)

	r, err := http.NewRequest("POST", trackingAPIURL, bytes.NewBufferString(trackingData))
	if err != nil {
		return b, err
	}

	r.Header.Add("Content-Type", "application/xml")

	resp, err := client.Client.Do(r)
	if err != nil {
		return b, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return b, err
	}

	return body, nil
}

func formatBody(body []byte) interface{} {
	var returnMap []interface{}

	objects := gjson.Get(string(body), "objeto")

	for _, objectValue := range objects.Array() {
		var object object

		object.Number = gjson.Get(objectValue.String(), "numero").String()
		object.Category = gjson.Get(objectValue.String(), "categoria").String()
		object.Error = returnError(object)

		if object.Error != "" {
			errorObject := map[string]string{
				"number": object.Number,
				"error":  object.Error,
			}
			returnMap = append(returnMap, errorObject)
			continue
		}

		events := gjson.Get(objectValue.String(), "evento")

		for key, event := range events.Array() {
			if key == 0 {
				putLastData(&object, event)
			} else {
				putHistory(&object, event)
			}
		}
		returnMap = append(returnMap, object)
	}

	return returnMap
}

func returnDestination(destination gjson.Result) string {
	local := gjson.Get(destination.String(), "local")
	place := gjson.Get(destination.String(), "endereco.logradouro")
	number := gjson.Get(destination.String(), "endereco.numero")
	locality := gjson.Get(destination.String(), "endereco.localidade")
	uf := gjson.Get(destination.String(), "endereco.uf")
	neighborhood := gjson.Get(destination.String(), "endereco.bairro")

	return fmt.Sprintf("%v - %v, %v, %v - %v/%v", local, place, number, neighborhood, locality, uf)
}

func returnOrigin(event gjson.Result) string {
	local := gjson.Get(event.String(), "unidade.local")
	place := gjson.Get(event.String(), "unidade.endereco.logradouro")
	number := gjson.Get(event.String(), "unidade.endereco.numero")
	locality := gjson.Get(event.String(), "unidade.endereco.localidade")
	uf := gjson.Get(event.String(), "unidade.endereco.uf")
	neighborhood := gjson.Get(event.String(), "unidade.endereco.bairro")

	return fmt.Sprintf("%v - %v, %v, %v - %v/%v", local, place, number, neighborhood, locality, uf)
}

func returnDate(event gjson.Result) string {
	date := gjson.Get(event.String(), "data")
	hour := gjson.Get(event.String(), "hora")

	return fmt.Sprintf("%v - %v", date, hour)
}

func returnError(obj object) string {
	if strings.Contains(obj.Category, "ERRO") {
		return obj.Category
	}

	return ""
}

func putLastData(obj *object, event gjson.Result) {
	obj.Date = returnDate(event)
	obj.Type = gjson.Get(event.String(), "tipo").String()
	obj.Status = int(gjson.Get(event.String(), "status").Int())
	obj.Description = gjson.Get(event.String(), "descricao").String()
	obj.Detail = gjson.Get(event.String(), "detalhe").String()
	obj.Origin = returnOrigin(event)

	destinations := gjson.Get(event.String(), "destino")
	for _, destination := range destinations.Array() {
		obj.Destination = returnDestination(destination)
	}
}

func putHistory(obj *object, event gjson.Result) {
	history := history{}

	history.Date = returnDate(event)
	history.Type = gjson.Get(event.String(), "tipo").String()
	history.Status = int(gjson.Get(event.String(), "status").Int())
	history.Description = gjson.Get(event.String(), "descricao").String()
	history.Detail = gjson.Get(event.String(), "detalhe").String()
	history.Origin = returnOrigin(event)

	destinations := gjson.Get(event.String(), "destino")
	for _, destination := range destinations.Array() {
		history.Destination = returnDestination(destination)
	}

	obj.History = append(obj.History, history)
}
