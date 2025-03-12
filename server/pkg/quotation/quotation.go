package quotation

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/getenv"
)

type QuotationResponse struct {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(getenv.QuotationApiTimeout))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getenv.QuotationApiEndpoint, nil)
	if err != nil {
		http.Error(w, "error while creating a request with context", http.StatusInternalServerError)
		return
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "error when making the request to the quotation api", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "error while reading the body response", http.StatusInternalServerError)
		return
	}

	var quotationResponse QuotationResponse
	err = json.Unmarshal(body, &quotationResponse)
	if err != nil {
		http.Error(w, "error while unmarshalling the body response into go struct", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(quotationResponse.Bid))
}
