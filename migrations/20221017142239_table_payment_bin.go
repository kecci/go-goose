package migrations

import (
	"database/sql"
	"log"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upTablePaymentBin, downTablePaymentBin)
}

func upTablePaymentBin(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	query := `
	CREATE TABLE IF NOT EXISTS public.payment_bin (
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

func downTablePaymentBin(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	query := `
	DROP TABLE public.payment_bin;
	`
	_, err := tx.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
