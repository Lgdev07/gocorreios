package tracking

var jsonResult = `{
    "objeto": [
        {
            "categoria": "SEDEX",
            "evento": [
                {
                    "cepDestino": "09937645",
                    "criacao": "05032021140435",
                    "data": "05/03/2021",
                    "dataPostagem": "03/03/2021",
                    "descricao": "Objeto entregue ao destinatário",
                    "diasUteis": "0",
                    "hora": "14:04",
                    "prazoGuarda": "0",
                    "recebedor": {},
                    "status": "01",
                    "tipo": "BDE",
                    "unidade": {
                        "cidade": "Diadema",
                        "codigo": "09951971",
                        "endereco": {
                            "bairro": "VILA NOGUEIRA",
                            "cep": "09951971",
                            "codigo": "34604",
                            "localidade": "Diadema",
                            "logradouro": "RUA ANTONIO DIAS ADORNO",
                            "numero": "236/240",
                            "uf": "SP"
                        },
                        "local": "CDD PIRAPORINHA",
                        "tipounidade": "Unidade de Distribuição",
                        "uf": "SP"
                    }
                },
                {
                    "cepDestino": "09937645",
                    "criacao": "05032021112458",
                    "data": "05/03/2021",
                    "dataPostagem": "03/03/2021",
                    "descricao": "Objeto saiu para entrega ao destinatário",
                    "diasUteis": "0",
                    "hora": "11:24",
                    "prazoGuarda": "0",
                    "status": "01",
                    "tipo": "OEC",
                    "unidade": {
                        "cidade": "Diadema",
                        "codigo": "09951971",
                        "endereco": {
                            "bairro": "VILA NOGUEIRA",
                            "cep": "09951971",
                            "localidade": "Diadema",
                            "logradouro": "RUA ANTONIO DIAS ADORNO",
                            "numero": "236/240",
                            "uf": "SP"
                        },
                        "local": "CDD PIRAPORINHA",
                        "tipounidade": "Unidade de Distribuição",
                        "uf": "SP"
                    }
                },
                {
                    "cepDestino": "09937645",
                    "criacao": "04032021100233",
                    "data": "04/03/2021",
                    "dataPostagem": "03/03/2021",
                    "descricao": "Área com distribuição sujeita a prazo diferenciado",
                    "detalhe": "Restrição de entrega domiciliar temporária",
                    "diasUteis": "0",
                    "hora": "10:02",
                    "prazoGuarda": "0",
                    "status": "08",
                    "tipo": "FC",
                    "unidade": {
                        "cidade": "Diadema",
                        "codigo": "09951971",
                        "endereco": {
                            "bairro": "VILA NOGUEIRA",
                            "cep": "09951971",
                            "codigo": "34604",
                            "localidade": "Diadema",
                            "logradouro": "RUA ANTONIO DIAS ADORNO",
                            "numero": "236/240",
                            "uf": "SP"
                        },
                        "local": "CDD PIRAPORINHA",
                        "tipounidade": "Unidade de Distribuição",
                        "uf": "SP"
                    }
                },
                {
                    "cepDestino": "09937645",
                    "criacao": "03032021223539",
                    "data": "03/03/2021",
                    "dataPostagem": "03/03/2021",
                    "descricao": "Objeto em trânsito - por favor aguarde",
                    "destino": [
                        {
                            "bairro": "VILA NOGUEIRA",
                            "cidade": "Diadema",
                            "codigo": "09951971",
                            "endereco": {
                                "bairro": "VILA NOGUEIRA",
                                "cep": "09951971",
                                "codigo": "34604",
                                "localidade": "Diadema",
                                "logradouro": "RUA ANTONIO DIAS ADORNO",
                                "numero": "236/240",
                                "uf": "SP"
                            },
                            "local": "CDD PIRAPORINHA",
                            "uf": "SP"
                        }
                    ],
                    "diasUteis": "0",
                    "hora": "22:35",
                    "prazoGuarda": "0",
                    "status": "01",
                    "tipo": "DO",
                    "unidade": {
                        "cidade": "SAO PAULO",
                        "codigo": "04293970",
                        "endereco": {
                            "bairro": "SAUDE",
                            "cep": "04293970",
                            "codigo": "41919",
                            "localidade": "SAO PAULO",
                            "logradouro": "RUA DO BOQUEIRAO",
                            "numero": "320",
                            "uf": "SP"
                        },
                        "local": "CTE SAUDE",
                        "tipounidade": "Unidade de Tratamento",
                        "uf": "SP"
                    }
                },
                {
                    "cepDestino": "09937645",
                    "criacao": "03032021143827",
                    "data": "03/03/2021",
                    "dataPostagem": "03/03/2021",
                    "descricao": "Objeto em trânsito - por favor aguarde",
                    "destino": [
                        {
                            "bairro": "SAUDE",
                            "cidade": "SAO PAULO",
                            "codigo": "04293970",
                            "endereco": {
                                "bairro": "SAUDE",
                                "cep": "04293970",
                                "codigo": "41919",
                                "localidade": "SAO PAULO",
                                "logradouro": "RUA DO BOQUEIRAO",
                                "numero": "320",
                                "uf": "SP"
                            },
                            "local": "CTE SAUDE",
                            "uf": "SP"
                        }
                    ],
                    "diasUteis": "0",
                    "hora": "14:38",
                    "prazoGuarda": "0",
                    "status": "01",
                    "tipo": "RO",
                    "unidade": {
                        "cidade": "SAO PAULO",
                        "codigo": "05777971",
                        "endereco": {
                            "bairro": "PIRAJUSSARA",
                            "cep": "05777970",
                            "codigo": "333751",
                            "complemento": "LOJA 109 A 110",
                            "localidade": "SAO PAULO",
                            "logradouro": "ESTRADA DO CAMPO LIMPO",
                            "numero": "459",
                            "uf": "SP"
                        },
                        "local": "AGF SHOPPING CAMPO LIMPO",
                        "tipounidade": "Agência dos Correios",
                        "uf": "SP"
                    }
                },
                {
                    "cepDestino": "09937645",
                    "criacao": "03032021140543",
                    "data": "03/03/2021",
                    "dataPostagem": "03/03/2021",
                    "descricao": "Objeto postado",
                    "diasUteis": "0",
                    "hora": "14:05",
                    "prazoGuarda": "0",
                    "status": "01",
                    "tipo": "PO",
                    "unidade": {
                        "cidade": "SAO PAULO",
                        "codigo": "05777971",
                        "endereco": {
                            "bairro": "PIRAJUSSARA",
                            "cep": "05777970",
                            "codigo": "333751",
                            "complemento": "LOJA 109 A 110",
                            "localidade": "SAO PAULO",
                            "logradouro": "ESTRADA DO CAMPO LIMPO",
                            "numero": "459",
                            "uf": "SP"
                        },
                        "local": "AGF SHOPPING CAMPO LIMPO",
                        "tipounidade": "Agência dos Correios",
                        "uf": "SP"
                    }
                }
            ],
            "nome": "ETIQUETA LOGICA SEDEX",
            "numero": "ON732904576BR",
            "sigla": "ON"
        }
    ],
    "pesquisa": "Lista de Objetos",
    "quantidade": "1",
    "resultado": "Todos os eventos",
    "versao": "3.0"
}`
