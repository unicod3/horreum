package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateProductsAndArticleTable, downCreateProductsAndArticleTable)
}

func upCreateProductsAndArticleTable(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`CREATE TABLE products (
    						id bigserial primary key,
    						created_at  timestamp without time zone DEFAULT now() NOT NULL,
    						updated_at  timestamp without time zone DEFAULT now() NOT NULL, 
    						sku varchar(256) not null,
    						price bigint  not null
						);`)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`CREATE TABLE articles (
    						id bigserial primary key,
    						created_at  timestamp without time zone DEFAULT now() NOT NULL,
    						updated_at  timestamp without time zone DEFAULT now() NOT NULL, 
    						sku varchar(256) not null,
    						quantity bigint  not null
						);`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`CREATE TABLE product_articles (
    						id bigserial primary key,
    						product_id bigint not null,
    						article_id bigint not null,
    						CONSTRAINT fk_products
									FOREIGN KEY(product_id) 
									REFERENCES products(id)
									ON DELETE CASCADE,
                              
    						CONSTRAINT fk_articles
									FOREIGN KEY(article_id) 
									REFERENCES articles(id)
									ON DELETE CASCADE
                              
						);`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateProductsAndArticleTable(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("DROP TABLE product_articles;")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE articles;")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE products;")
	if err != nil {
		return err
	}
	return nil
}
