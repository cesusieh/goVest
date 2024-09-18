package main

import (
	app "goVest/src"
	"log"
	"os"
)

func main() {
	app := app.Gerar()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
