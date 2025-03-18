package quotation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/getenv"
)

type QuotationResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
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

	json.NewEncoder(w).Encode(quotationResponse.USDBRL)
}
