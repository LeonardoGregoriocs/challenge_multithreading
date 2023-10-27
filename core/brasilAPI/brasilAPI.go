package brasilapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func GetAddressBrasilAPI(cep string, channelBrasilAPI chan BrasilAPI) (chan BrasilAPI, error) {
	// time.Sleep(time.Second * 1)
	resp, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != err {
		return nil, err
	}

	var c BrasilAPI
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	channelBrasilAPI <- c

	return channelBrasilAPI, nil
}
