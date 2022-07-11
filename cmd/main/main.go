package main

import (
	"log"

	"cmd/main/main.go/internal/app"
)

func main() {
	myApp := app.New()
	log.Fatal(myApp.Start())
}
