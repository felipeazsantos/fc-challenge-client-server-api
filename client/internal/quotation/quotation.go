package quotation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/felipeazsantos/fc-challenge-client-server-api/client/internal/dto"
	"github.com/felipeazsantos/fc-challenge-client-server-api/client/internal/getenv"
)

func MakeRequestOnServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(getenv.ClientTimeOut))
	defer cancel()

	url := fmt.Sprintf("%s/cotacao", getenv.ServerUrl)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("context timeout error: %v", err)
		}
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var quotation dto.QuotationDto
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		return err
	}

	// TODO: O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}	
	saveQuotationOnTxtFile(&quotation)

	return nil
}

func saveQuotationOnTxtFile(quotation *dto.QuotationDto) error {

	return nil
}