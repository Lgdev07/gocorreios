<h1 align="center">
      <img alt="go-correios" title="go-correios" src=".github/logo.png" width="300px" />
</h1>

<h3 align="center">
  Go Correios
</h3>

<p align="center">Integration with Correios Web Services✉️</p>
<p align="center">Made with Golang 🚀</p>

<p align="center">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/Lgdev07/go-correios?color=%2304D361">

  <img alt="Made by Lgdev07" src="https://img.shields.io/badge/made%20by-Lgdev07-%2304D361">

  <img alt="License" src="https://img.shields.io/badge/license-MIT-%2304D361">

  <a href="https://github.com/Lgdev07/go-correios/stargazers">
    <img alt="Stargazers" src="https://img.shields.io/github/stars/Lgdev07/go-correios?style=social">
  </a>
</p>

<p align="center">
  <a href="#-current-features">Features</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-use">How to Use</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-contribute">How to contribute</a>&nbsp;&nbsp;&nbsp;
</p>

## ✅ Features
- [x] Multiple Object Tracking
- [ ] CEP
- [ ] Fare Calculation

## 🚀 How to Use
## Tracking

```go
package main

import (
	"fmt"
	"log"

	"github.com/Lgdev07/go-correios/tracking"
)

func main() {
	result, err := tracking.Tracking([]string{"OK816158697BR"})
	if err != nil {
		log.Fatal(err)
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
        "last_status": "Objeto entregue ao remetente",
        "last_detail": "",
        "last_origin": "CDD CARIACICA - RUA SAO FRANCISCO, 53, SAO FRANCISCO - VIANA/ES",
        "last_destination": "",
        "history": [
            {
                "date": "05/10/2020 - 06:37",
                "status": "Objeto saiu para entrega ao remetente",
                "detail": "",
                "origin": "CDD CARIACICA - RUA SAO FRANCISCO, 53, SAO FRANCISCO - VIANA/ES",
                "destination": ""
            },
            {
                "date": "01/10/2020 - 22:07",
                "status": "Objeto encaminhado ",
                "detail": "",
                "origin": "CTE BENFICA - RUA LEOPOLDO BULHOES, 530, BENFICA - RIO DE JANEIRO/RJ",
                "destination": "CDD CARIACICA - RUA SAO FRANCISCO, 53, SAO FRANCISCO - VIANA/ES"
            },
            {
                "date": "30/09/2020 - 11:03",
                "status": "Objeto encaminhado ",
                "detail": "",
                "origin": "CEE CAMPOS DOS GOYTACAZES - RUA ROCHA LEAO, 120, PARQUE CAJU - CAMPOS DOS GOYTACAZES/RJ",
                "destination": "CTE BENFICA - RUA LEOPOLDO BULHOES, 530, BENFICA - RIO DE JANEIRO/RJ"
            },
            {
                "date": "30/09/2020 - 11:02",
                "status": "Objeto mal encaminhado",
                "detail": "Encaminhamento a ser corrigido",
                "origin": "CEE CAMPOS DOS GOYTACAZES - RUA ROCHA LEAO, 120, PARQUE CAJU - CAMPOS DOS GOYTACAZES/RJ",
                "destination": ""
            },
            {
                "date": "29/09/2020 - 18:06",
                "status": "Objeto encaminhado ",
                "detail": "",
                "origin": "CTE BENFICA - RUA LEOPOLDO BULHOES, 530, BENFICA - RIO DE JANEIRO/RJ",
                "destination": "CEE CAMPOS DOS GOYTACAZES - RUA ROCHA LEAO, 120, PARQUE CAJU - CAMPOS DOS GOYTACAZES/RJ"
            },
            {
                "date": "28/09/2020 - 21:03",
                "status": "Objeto encaminhado ",
                "detail": "",
                "origin": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP",
                "destination": "CTE BENFICA - RUA LEOPOLDO BULHOES, 530, BENFICA - RIO DE JANEIRO/RJ"
            },
            {
                "date": "28/09/2020 - 14:53",
                "status": "Objeto encaminhado ",
                "detail": "",
                "origin": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP",
                "destination": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP"
            },
            {
                "date": "26/09/2020 - 09:56",
                "status": "Destinatário não retirou objeto no prazo",
                "detail": "Objeto será devolvido ao remetente",
                "origin": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP",
                "destination": ""
            },
            {
                "date": "18/09/2020 - 09:42",
                "status": "Objeto aguardando retirada no endereço indicado",
                "detail": "Para retirá-lo, é preciso informar o código do objeto e apresentar documentação que comprove ser o destinatário ou pessoa por ele oficialmente autorizada.",
                "origin": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP",
                "destination": ""
            },
            {
                "date": "16/09/2020 - 20:57",
                "status": "Objeto encaminhado ",
                "detail": "",
                "origin": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP",
                "destination": "CDD PIRAPORINHA - RUA ANTONIO DIAS ADORNO, 236/240, VILA NOGUEIRA - Diadema/SP"
            },
            {
                "date": "16/09/2020 - 15:11",
                "status": "Objeto encaminhado ",
                "detail": "",
                "origin": "CTE JAGUARE - RUA MERGENTHALER, 568, VILA LEOPOLDINA - SAO PAULO/SP",
                "destination": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP"
            },
            {
                "date": "16/09/2020 - 01:53",
                "status": "Objeto encaminhado ",
                "detail": "",
                "origin": "AGF XV DE NOVEMBRO - AVENIDA QUINZE DE NOVEMBRO, 1668, CENTRO - ITAPECERICA DA SERRA/SP",
                "destination": "CTE JAGUARE - RUA MERGENTHALER, 568, VILA LEOPOLDINA - SAO PAULO/SP"
            },
            {
                "date": "16/09/2020 - 01:36",
                "status": "Objeto postado",
                "detail": "",
                "origin": "AGF XV DE NOVEMBRO - AVENIDA QUINZE DE NOVEMBRO, 1668, CENTRO - ITAPECERICA DA SERRA/SP",
                "destination": ""
            }
        ]
    }
]
```

## 🤔 How to contribute

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m 'feat: My new feature'`;
- Push to your branch: `git push origin my-feature`.

After the merge of your pull request is done, you can delete your branch.

---