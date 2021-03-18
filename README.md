<h1 align="center">
      <img alt="gocorreios" title="gocorreios" src=".github/logo.png" width="300px" />
</h1>

<h3 align="center">
  Go Correios
</h3>

<p align="center">Integration with Correios Web Services ✉️</p>
<p align="center">Made with Golang 🚀</p>

<p align="center">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/Lgdev07/gocorreios?color=%2304D361">

  <img alt="Made by Lgdev07" src="https://img.shields.io/badge/made%20by-Lgdev07-%2304D361">

  <img alt="License" src="https://img.shields.io/badge/license-MIT-%2304D361">

  <a href="https://github.com/Lgdev07/gocorreios/stargazers">
    <img alt="Stargazers" src="https://img.shields.io/github/stars/Lgdev07/gocorreios?style=social">
  </a>
</p>

<p align="center">
  <a href="#-current-features">Features</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-use">How to Use</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-contribute">How to contribute</a>&nbsp;&nbsp;&nbsp;
</p>

## ✅ Features
- [x] [Multiple Object Tracking](#tracking)
- [x] [CEP](#cep)
- [x] [Fare Calculation](#fare)

## 🚀 How to Use

<a id="tracking"></a>
### Tracking

```go
package main

import (
	"fmt"
	"log"

	"github.com/Lgdev07/gocorreios"
)

func main() {
	result, err := gocorreios.Tracking([]string{"OK816158697BR"})
	if err != nil {
        fmt.Println("[ERRO]:", err.Error())
        return
	}
	fmt.Println(string(result))
}


```

Expected result:

```
[
    {
        "number": "OK816158697BR",
        "category": "SEDEX",
        "last_date": "05/10/2020 - 09:09",
        "last_type": "BDE",
        "last_status": 23,
        "last_description": "Objeto entregue ao remetente",
        "last_detail": "",
        "last_origin": "CDD CARIACICA - RUA SAO FRANCISCO, 53, SAO FRANCISCO - VIANA/ES",
        "last_destination": "",
        "error": "",
        "history": [
            {
                "date": "05/10/2020 - 06:37",
                "type": "OEC",
                "status": 9,
                "description": "Objeto saiu para entrega ao remetente",
                "detail": "",
                "origin": "CDD CARIACICA - RUA SAO FRANCISCO, 53, SAO FRANCISCO - VIANA/ES",
                "destination": ""
            },
            {
                "date": "01/10/2020 - 22:07",
                "type": "DO",
                "status": 1,
                "description": "Objeto em trânsito - por favor aguarde",
                "detail": "",
                "origin": "CTE BENFICA - RUA LEOPOLDO BULHOES, 530, BENFICA - RIO DE JANEIRO/RJ",
                "destination": "CDD CARIACICA - RUA SAO FRANCISCO, 53, SAO FRANCISCO - VIANA/ES"
            },
            {
                "date": "30/09/2020 - 11:03",
                "type": "RO",
                "status": 1,
                "description": "Objeto em trânsito - por favor aguarde",
                "detail": "",
                "origin": "CEE CAMPOS DOS GOYTACAZES - RUA ROCHA LEAO, 120, PARQUE CAJU - CAMPOS DOS GOYTACAZES/RJ",
                "destination": "CTE BENFICA - RUA LEOPOLDO BULHOES, 530, BENFICA - RIO DE JANEIRO/RJ"
            },
            {
                "date": "30/09/2020 - 11:02",
                "type": "FC",
                "status": 3,
                "description": "Objeto mal encaminhado",
                "detail": "Encaminhamento a ser corrigido",
                "origin": "CEE CAMPOS DOS GOYTACAZES - RUA ROCHA LEAO, 120, PARQUE CAJU - CAMPOS DOS GOYTACAZES/RJ",
                "destination": ""
            },
            {
                "date": "29/09/2020 - 18:06",
                "type": "DO",
                "status": 1,
                "description": "Objeto em trânsito - por favor aguarde",
                "detail": "",
                "origin": "CTE BENFICA - RUA LEOPOLDO BULHOES, 530, BENFICA - RIO DE JANEIRO/RJ",
                "destination": "CEE CAMPOS DOS GOYTACAZES - RUA ROCHA LEAO, 120, PARQUE CAJU - CAMPOS DOS GOYTACAZES/RJ"
            },
            {
                "date": "28/09/2020 - 21:03",
                "type": "DO",
                "status": 1,
                "description": "Objeto em trânsito - por favor aguarde",
                "detail": "",
                "origin": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP",
                "destination": "CTE BENFICA - RUA LEOPOLDO BULHOES, 530, BENFICA - RIO DE JANEIRO/RJ"
            },
            {
                "date": "28/09/2020 - 14:53",
                "type": "RO",
                "status": 1,
                "description": "Objeto em trânsito - por favor aguarde",
                "detail": "",
                "origin": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP",
                "destination": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP"
            },
            {
                "date": "26/09/2020 - 09:56",
                "type": "BDI",
                "status": 26,
                "description": "Destinatário não retirou objeto no prazo",
                "detail": "Objeto será devolvido ao remetente",
                "origin": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP",
                "destination": ""
            },
            {
                "date": "18/09/2020 - 09:42",
                "type": "LDI",
                "status": 1,
                "description": "Objeto aguardando retirada no endereço indicado",
                "detail": "Para retirá-lo, é preciso informar o código do objeto e apresentar documentação que comprove ser o destinatário ou pessoa por ele oficialmente autorizada.",
                "origin": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP",
                "destination": ""
            },
            {
                "date": "16/09/2020 - 20:57",
                "type": "DO",
                "status": 1,
                "description": "Objeto em trânsito - por favor aguarde",
                "detail": "",
                "origin": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP",
                "destination": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP"
            },
            {
                "date": "16/09/2020 - 15:11",
                "type": "DO",
                "status": 1,
                "description": "Objeto em trânsito - por favor aguarde",
                "detail": "",
                "origin": "CTE JAGUARE - RUA MERGENTHALER, 568, VILA LEOPOLDINA - SAO PAULO/SP",
                "destination": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP"
            },
            {
                "date": "16/09/2020 - 01:53",
                "type": "RO",
                "status": 1,
                "description": "Objeto em trânsito - por favor aguarde",
                "detail": "",
                "origin": "AGF XV DE NOVEMBRO - AVENIDA QUINZE DE NOVEMBRO, 1668, CENTRO - ITAPECERICA DA SERRA/SP",
                "destination": "CTE JAGUARE - RUA MERGENTHALER, 568, VILA LEOPOLDINA - SAO PAULO/SP"
            },
            {
                "date": "16/09/2020 - 01:36",
                "type": "PO",
                "status": 1,
                "description": "Objeto postado",
                "detail": "",
                "origin": "AGF XV DE NOVEMBRO - AVENIDA QUINZE DE NOVEMBRO, 1668, CENTRO - ITAPECERICA DA SERRA/SP",
                "destination": ""
            }
        ]
    }
]
```
<a id="cep"></a>
### CEP

```go
package main

import (
	"fmt"

	"github.com/Lgdev07/gocorreios"
)

func main() {
	res, err := gocorreios.Cep("04029000")
	if err != nil {
		fmt.Println("[ERRO]:", err.Error())
		return
	}

	fmt.Println(string(res))
}
```

Expected result:
```
{
    "cep": "04029-000",
    "uf": "SP",
    "localidade": "São Paulo",
    "logradouro": "Avenida Ibirapuera",
    "bairro": "Indianópolis",
    "complemento": "até 1731 - lado ímpar"
}
```

<a id="fare"></a>
### Fare

```go
package main

import (
	"fmt"

	"github.com/Lgdev07/gocorreios"
	"github.com/Lgdev07/gocorreios/fare"
)

func main() {
	params := fare.Interface{
		Service:        "SEDEX", // Avaliable Services: SEDEX, PAC
		CepOrigin:      "60010100",
		CepDestination: "04029000",
		Weight:         30,
		Length:         20,
		Height:         20,
		Width:          20,
	}

	result, err := gocorreios.Fare(params)
	if err != nil {
		fmt.Println("[ERRO]:", err.Error())
		return
	}
	fmt.Println(string(result))
}
```

Expected result:
```
{
    "service": "SEDEX",
    "price": "22,50",
    "days_for_delivery": "9",
    "deliver_home": "S",
    "deliver_saturday": "S",
    "obs": "O CEP de destino está sujeito a condições especiais de entrega  pela  ECT e será realizada com o acréscimo de até 7 (sete) dias úteis ao prazo regular."
}
```


## 🤔 How to contribute

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m 'feat: My new feature'`;
- Push to your branch: `git push origin my-feature`.

After the merge of your pull request is done, you can delete your branch.

---
