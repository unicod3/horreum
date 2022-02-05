package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateWarehousesTable, downCreateWarehousesTable)
}

func upCreateWarehousesTable(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`CREATE TABLE warehouses (
    						id bigserial primary key,
    						created_at timestamp with time zone not null,
    						updated_at timestamp with time zone not null, 
    						name varchar(256) not null
						);`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateWarehousesTable(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE warehouses;")
	if err != nil {
		return err
	}
	return nil
}
