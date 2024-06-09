package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, postcode := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + postcode + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Request fail: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error to read the response: %v\n", err)
		}

		var data ViaCEP

		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error to parse the response: %v\n", err)

		}

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error to create the file: %v\n", err)
		}

		defer file.Close()

		fileContentFormatted := fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf)
		_, err = file.WriteString(fileContentFormatted)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error to write the file: %v\n", err)
		}

		fmt.Println("File created with success.")
		fmt.Println("City: ", data.Localidade)

	}
}
