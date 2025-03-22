package quotation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	err = saveQuotationOnTxtFile(&quotation)
	if err != nil {
		return err
	}

	return nil
}

func saveQuotationOnTxtFile(quotation *dto.QuotationDto) error {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", quotation.Bid))
	if err != nil {
		return err
	}

	return nil
}