package quotation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/getenv"
	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/model"
	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/repository"
)

type QuotationResponse struct {
	USDBRL model.USDBRL `json:"USDBRL"`
}



func GetUSDBRLQuotation(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(getenv.QuotationApiTimeout*uint64(time.Millisecond)))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getenv.QuotationApiEndpoint, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while creating a request with context: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when making the request to the quotation api: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while reading the body response: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	var quotationResponse QuotationResponse
	err = json.Unmarshal(body, &quotationResponse)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while unmarshalling the body response into go struct: %s", err.Error()), http.StatusInternalServerError)
		return
	}


	err = repository.QuotationRepository.InsertQuotation(&quotationResponse.USDBRL)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while inserting quotation into database: %s", err.Error()), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotationResponse.USDBRL)
}
