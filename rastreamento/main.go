package rastreamento

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// Objeto representa a estrutura principal
type Objeto struct {
	Numero    string      `json:"numero"`
	Categoria string      `json:"categoria"`
	Data      string      `json:"ultima_data"`
	Situacao  string      `json:"ultima_situacao"`
	Detalhe   string      `json:"ultimo_detalhe"`
	Origem    string      `json:"ultima_origem"`
	Destino   string      `json:"ultimo_destino"`
	Historico []Historico `json:"historico"`
}

// Historico representa a lista de eventos relacionados a um Objeto
type Historico struct {
	Data     string `json:"data"`
	Situacao string `json:"situacao"`
	Detalhe  string `json:"detalhe"`
	Origem   string `json:"origem"`
	Destino  string `json:"destino"`
}

func Rastreamento() []byte {
	_, body := BuscaCodigo("OK816158697BR")

	valores := []Objeto{}

	objetos := gjson.Get(string(body), "objeto")

	for _, objetoValue := range objetos.Array() {
		objeto := Objeto{}
		objeto.Numero = gjson.Get(objetoValue.String(), "numero").String()
		objeto.Categoria = gjson.Get(objetoValue.String(), "categoria").String()

		eventos := gjson.Get(objetoValue.String(), "evento")
		for key, evento := range eventos.Array() {
			if key == 0 {

				objeto.Data = BuscaData(evento)
				objeto.Situacao = gjson.Get(evento.String(), "descricao").String()
				objeto.Detalhe = gjson.Get(evento.String(), "detalhe").String()
				objeto.Origem = BuscaOrigem(evento)

				destinos := gjson.Get(evento.String(), "destino")
				for _, destino := range destinos.Array() {
					objeto.Destino = BuscaDestino(destino)
				}

			} else {
				historico := Historico{}

				historico.Data = BuscaData(evento)
				historico.Situacao = gjson.Get(evento.String(), "descricao").String()
				historico.Detalhe = gjson.Get(evento.String(), "detalhe").String()
				historico.Origem = BuscaOrigem(evento)

				destinos := gjson.Get(evento.String(), "destino")
				for _, destino := range destinos.Array() {
					historico.Destino = BuscaDestino(destino)
				}

				objeto.Historico = append(objeto.Historico, historico)
			}

		}
		valores = append(valores, objeto)
	}

	e, err := json.Marshal(valores)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	return e

}

// BuscaDestino retorna o endereço do Destino
func BuscaDestino(destino gjson.Result) string {
	local := gjson.Get(destino.String(), "local")
	logradouro := gjson.Get(destino.String(), "endereco.logradouro")
	numero := gjson.Get(destino.String(), "endereco.numero")
	localidade := gjson.Get(destino.String(), "endereco.localidade")
	uf := gjson.Get(destino.String(), "endereco.uf")
	bairro := gjson.Get(destino.String(), "endereco.bairro")

	return fmt.Sprintf("%v - %v, %v, %v - %v/%v", local, logradouro, numero, localidade, uf, bairro)
}

// BuscaOrigem retorna o endereço da Origem
func BuscaOrigem(evento gjson.Result) string {
	local := gjson.Get(evento.String(), "unidade.local")
	logradouro := gjson.Get(evento.String(), "unidade.endereco.logradouro")
	numero := gjson.Get(evento.String(), "unidade.endereco.numero")
	localidade := gjson.Get(evento.String(), "unidade.endereco.localidade")
	uf := gjson.Get(evento.String(), "unidade.endereco.uf")
	bairro := gjson.Get(evento.String(), "unidade.endereco.bairro")

	return fmt.Sprintf("%v - %v, %v, %v - %v/%v", local, logradouro, numero, bairro, localidade, uf)
}

// BuscaData retorna a data do evento
func BuscaData(evento gjson.Result) string {
	data := gjson.Get(evento.String(), "data")
	hora := gjson.Get(evento.String(), "hora")
	return fmt.Sprintf("%v - %v", data, hora)
}

func BuscaCodigo(_code string) (string, []byte) {
	apiUrl := "http://webservice.correios.com.br/service/rest/rastro/rastroMobile"
	data := "<rastroObjeto><usuario></usuario><senha></senha><tipo>L</tipo><resultado>T</resultado><objetos>" + _code + "</objetos><lingua>101</lingua><token></token></rastroObjeto>"
	client := &http.Client{}

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

	return resp.Status, _body
}
