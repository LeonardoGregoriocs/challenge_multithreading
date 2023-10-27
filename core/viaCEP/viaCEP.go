package viacep

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"Ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func GetAddressViaCep(cep string, chanViaCep chan ViaCep) (chan ViaCep, error) {
	// time.Sleep(time.Second * 1)
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != err {
		return nil, err
	}

	var c ViaCep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	chanViaCep <- c

	return chanViaCep, nil
}
