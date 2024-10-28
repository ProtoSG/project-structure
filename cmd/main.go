package main

import (
	"log"
	"struct-proyect/cmd/api"
)

func main() {
	app := api.NewApiServer(":8080", nil)
	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
