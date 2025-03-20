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

	stmt, err := db.Prepare(`INSERT INTO quotation()`)


	return nil
}

