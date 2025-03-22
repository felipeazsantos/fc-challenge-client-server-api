package getenv

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ClientTimeOut uint64
	ServerUrl string
)

func LoadConfig() (err error) {
	if err = godotenv.Load(); err != nil {
		return
	}

	ClientTimeOut, err = strconv.ParseUint(os.Getenv("CLIENT_TIMEOUT"), 10, 64)
	if err != nil {
		return
	}

	ServerUrl = fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))

	return
}
