package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Resultado struct {
	Objetos []Objetos `json:"objeto"`
}

type Objetos struct {
	Numero  string   `json:"numero"`
	Eventos []Evento `json:"evento"`
}

type Evento struct {
	Data      string    `json:"data"`
	Hora      string    `json:"hora"`
	Descricao string    `json:"descricao"`
	Origem    Origem    `json:"unidade"`
	Destino   []Destino `json:"destino"`
}

type Origem struct {
	Local          string         `json:"local"`
	Cidade         string         `json:"cidade"`
	Uf             string         `json:"uf"`
	EnderecoOrigem EnderecoOrigem `json:"endereco"`
}

type EnderecoOrigem struct {
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Uf         string `json:"uf"`
	Bairro     string `json:"bairro"`
}

type Destino struct {
	Local  string `json:"local"`
	Cidade string `json:"cidade"`
	Uf     string `json:"uf"`
}

type EnderecoDestino struct {
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Uf         string `json:"uf"`
	Bairro     string `json:"bairro"`
}

func main() {
	_, body := BuscaCep("OK816158697BROK816158697BR")

	for _, objeto := range body.Objetos {
		println(objeto.Numero)

		for _, evento := range objeto.Eventos {
			println(evento.Data)
			println(evento.Descricao)
			println(evento.Hora)

			if len(evento.Destino) > 0 {
				println(evento.Destino[0].Cidade)
			}

		}
	}
}

func BuscaCep(_code string) (string, *Resultado) {
	apiUrl := "http://webservice.correios.com.br/service/rest/rastro/rastroMobile"
	data := "<rastroObjeto><usuario></usuario><senha></senha><tipo>L</tipo><resultado>T</resultado><objetos>" + _code + "</objetos><lingua>101</lingua><token></token></rastroObjeto>"
	client := &http.Client{}
	resultado := &Resultado{}

	r, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/xml")

	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(_body, &resultado)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Status, resultado
}

// [
// 	{
// 		"code": numero,
// 		"last_update": data - hora,
// 		"situation"
// 	}
// ]
