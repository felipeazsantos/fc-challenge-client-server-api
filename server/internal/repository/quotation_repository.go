package repository

import (
	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/database"
	"github.com/felipeazsantos/fc-challenge-client-server-api/server/internal/model"
)

type QuotationRepository struct {}


func (q *QuotationRepository) InsertQuotation(quotation *model.USDBRL) error {
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

	_, err = stmt.Exec(
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

