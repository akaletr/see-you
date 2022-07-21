package main

import (
	"log"

	"cmd/main/main.go/internal/app"
	"cmd/main/main.go/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	myApp := app.New(cfg)

	log.Fatal(myApp.Start())
}
