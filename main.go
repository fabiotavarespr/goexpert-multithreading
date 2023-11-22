package main

import (
	"encoding/json"
	"fmt"
	"github.com/fabiotavarespr/goexpert-multithreading/dto"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	urlViaCep       = "https://viacep.com.br/ws/%s/json/"
	urlBrasilApi    = "https://brasilapi.com.br/api/cep/v1/%s"
	urlBrasilAberto = "https://brasilaberto.com/api/v1/zipcode/%s"
	requestTimeout  = 1
)

func main() {
	cep := os.Args[1:]
	if len(cep) == 0 {
		log.Fatalf("Enter a zip code")
	}

	apiViaCep := make(chan dto.ViaCep)
	apiBrasilApi := make(chan dto.BrasilApi)
	apiBrasilAberto := make(chan dto.BrasilAberto)

	var viaCep dto.ViaCep
	var BrasilApi dto.BrasilApi
	var BrasilAberto dto.BrasilAberto
	go requestApi(fmt.Sprintf(urlViaCep, cep[0]), viaCep, apiViaCep)
	go requestApi(fmt.Sprintf(urlBrasilApi, cep[0]), BrasilApi, apiBrasilApi)
	go requestApi(fmt.Sprintf(urlBrasilAberto, cep[0]), BrasilAberto, apiBrasilAberto)

	select {
	case result1 := <-apiViaCep:
		formatResult("viacep.com.br", result1)
	case result2 := <-apiBrasilApi:
		formatResult("brasilapi.com.br", result2)
	case result3 := <-apiBrasilAberto:
		formatResult("brasilaberto.com", result3)
	case <-time.After(time.Second * requestTimeout):
		log.Fatalf("Timeout")
	}
}

func formatResult(provider string, result interface{}) {
	log.Printf("Received from %s: \n %v\n", provider, result)
}

func requestApi[T any](url string, apiStruct T, result chan T) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in request: %v", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}
	err = json.Unmarshal(body, &apiStruct)
	if err != nil {
		log.Fatalf("Error parsing json: %v", err)
	}
	result <- apiStruct
}
