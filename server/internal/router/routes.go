package router

import (
	"fmt"
	"net/http"

	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/getenv"
	"github.com/felipeazsantos/fc-challenge-client-server-api/server/pkg/quotation"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) Run() error {
	mux := s.configureRoutes()
	fmt.Println("Server running on port", getenv.ServerPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", getenv.ServerPort), mux)
}

func (s *server) configureRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", quotation.GetUSDBRLQuotation)
	return mux
}
