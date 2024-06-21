package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const ApiEndpoint = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

type ExchangeRateFromUsdToBrl struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

// O server.go deverá consumir a API contendo o câmbio de Dólar e Real
// no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL
// e em seguida deverá retornar no formato JSON o resultado para o cliente.

// Usando o package "context", o server.go deverá registrar
// no banco de dados SQLite cada cotação recebida,
// sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms
// e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	exchangeRate, err := GetExchangeRate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchangeRate)
}

func GetExchangeRate() (*ExchangeRateFromUsdToBrl, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ApiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	var exchangeRate ExchangeRateFromUsdToBrl
	err = json.Unmarshal(body, &exchangeRate)
	if err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}
