package main

import (
	"log"

	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/database"
	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/getenv"
	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/router"
)

func main() {
	if err := getenv.LoadConfig(); err != nil {
		log.Fatal("unable to load application configs", err)
	}

	if err := database.InitDB(); err != nil {
		log.Fatal("unable to load database", err)
	}

	svr := router.NewServer()
	log.Fatal(svr.Run())
}
