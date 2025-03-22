package getenv

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	QuotationApiEndpoint string
	QuotationApiTimeout  uint64
	DatabaseTimeout      uint64
	ServerPort           uint64
	ServerUrl			 string
)

func LoadConfig() error {
	var err error

	if err = godotenv.Load(); err != nil {
		return err
	}

	QuotationApiEndpoint = os.Getenv("QUOTATION_API_ENDPOINT")
	QuotationApiTimeout, err = strconv.ParseUint(os.Getenv("QUOTATION_API_TIMEOUT"), 10, 64)
	if err != nil {
		return err
	}

	DatabaseTimeout, err = strconv.ParseUint(os.Getenv("DATABASE_TIMEOUT"), 10, 64)
	if err != nil {
		return err
	}

	ServerPort, err = strconv.ParseUint(os.Getenv("SERVER_PORT"), 10, 64)
	if err != nil {
		return err
	}

	ServerUrl = fmt.Sprintf("%s:%d", os.Getenv("SERVER_HOST"), ServerPort)

	return nil
}
