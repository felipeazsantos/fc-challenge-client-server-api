package repository

import (
	"context"
	"time"

	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/database"
	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/getenv"
	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/model"
)

var QuotationRepository *quotationRepository = &quotationRepository{}


type quotationRepository struct {}


func (q *quotationRepository) InsertQuotation(quotation *model.USDBRL) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`INSERT INTO quotation(code, code_in, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date)
							 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}

	defer func() error {
		err = stmt.Close()
		if err != nil {
			return err
		}
		return nil
	}()


	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(getenv.DatabaseTimeout * uint64(time.Millisecond)))
	defer cancel()

	_, err = stmt.ExecContext(
		ctx,
		quotation.Code,
		quotation.Codein,
		quotation.Name,
		quotation.High,
		quotation.Low,
		quotation.VarBid,
		quotation.PctChange,
		quotation.Bid,
		quotation.Ask,
		quotation.Timestamp,
		quotation.CreateDate,
	)
	if err != nil {
		return err
	}

	return nil
}

func (q *quotationRepository) GetLastQuotation() (quotation *model.USDBRL, err error) {
	quotation = &model.USDBRL{}
	query := `
		SELECT code, code_in, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date
		  FROM quotation
		  ORDER BY create_date DESC
		  LIMIT 1
	`

	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(query).Scan(
		&quotation.Code,
		&quotation.Codein,
		&quotation.Name,
		&quotation.High,
		&quotation.Low,
		&quotation.VarBid,
		&quotation.PctChange,
		&quotation.Bid,
		&quotation.Ask,
		&quotation.Timestamp,
		&quotation.CreateDate,
	)
	if err != nil {
		return nil, err
	}

	return
}