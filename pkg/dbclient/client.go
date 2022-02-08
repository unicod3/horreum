package dbclient

import (
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

// Client holds database session
type Client struct {
	Session *db.Session
}

// DataStorage serves a contract over Client
type DataStorage interface {
	NewDataCollection(tableName string) DataTable
}

// DataCollection implements DataTable interface
type DataCollection struct {
	db.Collection
}

// DataTable serves a contract for DataCollection
type DataTable interface {
	db.Collection
	FindAll(dataAddress interface{}) error
	FindOne(cond Condition, dataAddress interface{}) error
	FindRelated(tableName string, condition Condition, dataAddress interface{}) error
	CreateRelated(tableName string, dataAddress interface{}) error
	Delete(cond Condition) error
	DeleteRelated(tableName string, condition Condition) error
	LoadMany2Many(columns, from, join, on string, condition Condition, dataAddress interface{}) error
}

// Condition is map to define query conditions
type Condition = db.Cond

// NewPostgresClient returns a Client struct which holds a postgres session
func NewPostgresClient(host, user, database, password string) DataStorage {
	settings := postgresql.ConnectionURL{
		Database: database,
		Host:     host,
		User:     user,
		Password: password,
	}
	db.LC().SetLevel(db.LogLevelInfo)
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	return &Client{Session: &sess}
}

// NewDataCollection returns a DataTable interface
func (client *Client) NewDataCollection(tableName string) DataTable {
	return &DataCollection{
		(*(client.Session)).Collection(tableName),
	}
}

// FindAll gets all the records for given DataTable
// and write it to given address
func (c *DataCollection) FindAll(dataAddress interface{}) error {
	if err := c.Find().All(dataAddress); err != nil {
		return err
	}
	return nil
}

// FindOne gets one record that matches the given Condition
// and writes it to given address
func (c *DataCollection) FindOne(cond Condition, dataAddress interface{}) error {
	if err := c.Find(cond).One(dataAddress); err != nil {
		return err
	}
	return nil
}

// Delete gets the records that matches the given Condition
// and deletes them
func (c *DataCollection) Delete(cond Condition) error {
	if err := c.Find(cond).Delete(); err != nil {
		return err
	}
	return nil
}

// FindRelated finds the related records of given table
func (c *DataCollection) FindRelated(tableName string, condition Condition, dataAddress interface{}) error {
	return c.Session().
		Collection(tableName).
		Find(condition).
		All(dataAddress)
}

// DeleteRelated deletes related records from given table
func (c *DataCollection) DeleteRelated(tableName string, condition Condition) error {
	return c.Session().
		Collection(tableName).
		Find(condition).
		Delete()
}

// CreateRelated creates record on the given related table
func (c *DataCollection) CreateRelated(tableName string, dataAddress interface{}) error {
	return c.Session().
		Collection(tableName).
		InsertReturning(dataAddress)
}

func (c *DataCollection) LoadMany2Many(columns, from, join, on string, condition Condition, dataAddress interface{}) error {
	return c.Session().SQL().
		Select(db.Raw(columns)).From(from).
		Join(join).On(on).
		Where(condition).
		All(dataAddress)
}
