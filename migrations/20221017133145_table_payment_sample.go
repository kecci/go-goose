package migrations

import (
	"database/sql"
	"log"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upTablePaymentSample, downTablePaymentSample)
}

func upTablePaymentSample(tx *sql.Tx) error {

	// This code is executed when the migration is applied.
	query := `
	CREATE TABLE IF NOT EXISTS public.payment_transactions_sample (
		id bigserial NOT NULL
	);
	`

	_, err := tx.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func downTablePaymentSample(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	query := `
	DROP TABLE public.payment_transactions_sample;
	`
	_, err := tx.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
