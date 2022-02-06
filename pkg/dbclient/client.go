package dbclient

import (
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type Client struct {
	Session *db.Session
}

// DataStore abstract away the data collection
type DataStore struct {
	db.Collection
}

// NewPostgresClient returns a Client struct which holds a postgres session
func NewPostgresClient(host, user, database, password string) *Client {
	settings := postgresql.ConnectionURL{
		Database: database,
		Host:     host,
		User:     user,
		Password: password,
	}
	db.LC().SetLevel(db.LogLevelTrace)
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	return &Client{Session: &sess}
}

func (client *Client) NewDataStore(tableName string) *DataStore {
	return &DataStore{
		(*(client.Session)).Collection(tableName),
	}
}
