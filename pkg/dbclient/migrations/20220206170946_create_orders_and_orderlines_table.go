package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateOrdersAndOrderlinesTable, downCreateOrdersAndOrderlinesTable)
}

func upCreateOrdersAndOrderlinesTable(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`CREATE TABLE orders (
    						id bigserial primary key,
    						warehouse_id bigint not null,
    						created_at  timestamp without time zone DEFAULT now() NOT NULL,
    						updated_at  timestamp without time zone DEFAULT now() NOT NULL, 
    						customer varchar(256) not null,
                        		
    						    CONSTRAINT fk_warehouse
									FOREIGN KEY(warehouse_id) 
									REFERENCES warehouses(id)
									ON DELETE CASCADE
						);`)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`CREATE TABLE order_lines (
    						id bigserial primary key,
    						order_id bigint not null,
    						created_at  timestamp without time zone DEFAULT now() NOT NULL,
    						updated_at  timestamp without time zone DEFAULT now() NOT NULL, 
    						sku varchar(256) not null,
    						quantity bigint not null,
    						unit_cost bigint  not null,
    						
    						   CONSTRAINT fk_order
									FOREIGN KEY(order_id) 
									REFERENCES orders(id)
									ON DELETE CASCADE
						);`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateOrdersAndOrderlinesTable(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("DROP TABLE order_lines;")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE orders;")
	if err != nil {
		return err
	}
	return nil
}
