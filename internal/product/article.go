package product

import (
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/streamer"
	"math"
	"time"
)

// ArticleRepository serves as a contract over ArticleService
type ArticleRepository interface {
	GetAll() ([]Article, error)
	GetById(uint64) (*Article, error)
	Create(*Article) error
	Update(*Article) error
	Delete(*Article) error
}

// Article represents a record from articles table
type Article struct {
	ID                 uint64    `json:"id" uri:"id" db:"id,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	Name               string    `json:"name" db:"name"`
	Stock              int64     `json:"stock" db:"stock"`
	AmountOf           int64     `json:"amount_of,omitempty" db:"amount_of,omitempty"`
	AvailableInventory int64     `json:"available_inventory,omitempty" db:"-"`
}

func (a *Article) CalculateAvailableInventory() {
	if a.AmountOf == 0 {
		a.AvailableInventory = 0
		return
	}
	a.AvailableInventory = int64(math.Floor(float64(a.Stock / a.AmountOf)))
}

// ArticleRequestBody represents the data type that needs to be sent over request
type ArticleRequestBody struct {
	Name  string `json:"name" db:"name"`
	Stock int64  `json:"stock" db:"stock"`
}

// ArticleService holds information about the datatable
// and implements ArticleService
type ArticleService struct {
	DataTable     dbclient.DataTable
	StreamChannel streamer.Channel
	StreamTopic   string
}

// GetAll returns all the records
func (service *ArticleService) GetAll() ([]Article, error) {
	var articles []Article

	if err := service.DataTable.FindAll(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// GetById returns single record for given pk id
func (service *ArticleService) GetById(id uint64) (*Article, error) {
	var article Article
	if err := service.DataTable.FindOne(dbclient.Condition{"id": id}, &article); err != nil {
		return nil, err
	}
	return &article, nil
}

// Create creates a new record on the datastore with given struct
func (service *ArticleService) Create(a *Article) error {
	if err := service.DataTable.InsertReturning(a); err != nil {
		return err
	}
	return nil
}

// Update updates given record on the datastore by finding it with its pk
func (service *ArticleService) Update(a *Article) error {
	a.UpdatedAt = time.Now().UTC()
	if err := service.DataTable.UpdateReturning(a); err != nil {
		return err
	}
	return nil
}

// Delete deletes the given struct from database by finding it with its pk
func (service *ArticleService) Delete(a *Article) error {
	if err := service.DataTable.Delete(dbclient.Condition{"id": a.ID}); err != nil {
		return err
	}
	return nil
}
