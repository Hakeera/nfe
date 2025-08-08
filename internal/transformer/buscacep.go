// Package transformer
package transformer 

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Endereco struct {
    CEP         string `json:"cep"`
    Logradouro  string `json:"logradouro"`
    Bairro      string `json:"bairro"`
    Localidade  string `json:"localidade"`
    UF          string `json:"uf"`
}

func GetEnderecoPorCEP(cep string) (*Endereco, error) {
    resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var endereco Endereco
    if err := json.NewDecoder(resp.Body).Decode(&endereco); err != nil {
        return nil, err
    }

    return &endereco, nil
}
