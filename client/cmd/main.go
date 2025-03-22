package main

import (
	"log"

	"github.com/felipeazsantos/fc-challenge-client-server-api/client/internal/getenv"
	"github.com/felipeazsantos/fc-challenge-client-server-api/client/internal/quotation"
)

func main() {
	if err := getenv.LoadConfig(); err != nil {
		log.Fatalf("error while trying load config: %v", err)
	}

	err := quotation.MakeRequestOnServer()
	if err != nil {
		log.Fatalf("error when trying making quotation request on server: %v", err)
	}

	log.Println("quotation request was successfully")
}
