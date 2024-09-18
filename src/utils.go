package app

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
)

type key struct {
	Key string
}

type stock struct {
	Symbol             string
	LongName           string
	RegularMarketPrice float64
	EarningsPerShare   float64
	PriceEarnings      float64
}

type result struct {
	Results []stock `json:"results"`
}

func getKey() string {
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatal("\nNão foi possível abrir o arquivo de configuração\n", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var key key
	decoder.Decode(&key)
	return key.Key
}

func getStock(r []byte) stock {
	var result result
	if err := json.Unmarshal(r, &result); err != nil {
		log.Fatal(err)
	}
	return result.Results[0]
}

func viewStock(r stock) {
	fmt.Printf(
		"Ativo: %s\nEmpresa: %s\nValor atual: %.2f\nValor de Graham: %.2f",
		r.Symbol, r.LongName, r.RegularMarketPrice, math.Sqrt(22.5*r.EarningsPerShare*r.PriceEarnings),
	)
}

func testKey(key string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://brapi.dev/api/quote/petr4?token=%s", key))
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
