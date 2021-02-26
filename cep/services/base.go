package services

// Item represents the cep structure
type Item struct {
	Cep          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	Complement   string `json:"complemento"`
}

// ResultError represents the result of cep response
type ResultError struct {
	Res Item
	Err error
}
