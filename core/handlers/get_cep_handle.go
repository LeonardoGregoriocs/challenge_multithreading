package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	brasilapi "github.com/leonardogregoriocs/challenge_multithreading/core/brasilAPI"
	viacep "github.com/leonardogregoriocs/challenge_multithreading/core/viaCEP"
)

func SearchZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channelviaCep := make(chan viacep.ViaCep)
	channelBrasilAPI := make(chan brasilapi.BrasilAPI)

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusConflict)
		return
	}

	cepParams := r.URL.Query().Get("cep")
	if cepParams == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go viacep.GetAddressViaCep(cepParams, channelviaCep)

	go brasilapi.GetAddressBrasilAPI(cepParams, channelBrasilAPI)

	select {
	case address := <-channelviaCep:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(address)
		fmt.Printf("Address: %v\nAPI: ViaCep\n", address)
	case address := <-channelBrasilAPI:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(address)
		fmt.Printf("Address: %v\nAPI: BrasilAPI\n", address)
	case <-time.After(time.Second * 1):
		w.WriteHeader(http.StatusRequestTimeout)
		json.NewEncoder(w).Encode("Err: Timeout")
		fmt.Println("Timeout")
	}
}
