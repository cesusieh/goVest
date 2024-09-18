package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func buscarAcao(c *cli.Context) {
	stock := c.String("stock")
	key := getKey()

	resp, err := http.Get(fmt.Sprintf("https://brapi.dev/api/quote/%s?token=%s", stock, key))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 404 {
		fmt.Println("Ativo não encontrado")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := getStock(body)
	viewStock(result)
}

func viewKey(c *cli.Context) {
	key := getKey()
	fmt.Println(key)
}

func registerKey(c *cli.Context) {
	if c.String("key") == "" {
		fmt.Println("Por favor, informe sua chave corretamente")
		return
	}
	key := key{
		Key: c.String("key"),
	}
	resultKey, err := testKey(key.Key)
	if err != nil {
		log.Fatal("\nOcorreu um erro ao testar sua chave.\n", err)
		return
	}
	if resultKey == 400 {
		fmt.Println("Sua chave de API parece ser inválida.")
		return
	}

	jsonKey, err := json.Marshal(key)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("conf.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write(jsonKey)
	fmt.Println("Chave de API cadastrada com sucesso!")
}
