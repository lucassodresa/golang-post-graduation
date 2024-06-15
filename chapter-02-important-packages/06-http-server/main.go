package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", GetAddressByPostcodeHandler)
	http.ListenAndServe(":8080", nil)
}

func GetAddressByPostcodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	postcodeParam := r.URL.Query().Get("postcode")
	if postcodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	address, error := GetAddressByPostcode(postcodeParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// result, error := json.Marshal(address)
	// if error != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)

	json.NewEncoder(w).Encode(address)
}

func GetAddressByPostcode(postcode string) (*ViaCEP, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + postcode + "/json")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var address ViaCEP
	error = json.Unmarshal(body, &address)
	if error != nil {
		return nil, error
	}
	return &address, nil
}
