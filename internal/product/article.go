package product

import (
	"github.com/unicod3/horreum/pkg/dbclient"
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
	ID        uint64    `json:"id" uri:"id" db:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	SKU       string    `json:"sku" db:"sku"`
	Quantity  int64     `json:"quantity" db:"quantity"`
}

// ArticleRequestBody represents the data type that needs to be sent over request
type ArticleRequestBody struct {
	SKU      string `json:"sku"`
	Quantity int64  `json:"quantity"`
}

// ArticleService holds information about the datatable
// and implements ArticleService
type ArticleService struct {
	dataTable dbclient.DataTable
}

// GetAll returns all the records
func (service *ArticleService) GetAll() ([]Article, error) {
	var articles []Article

	if err := service.dataTable.FindAll(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// GetById returns single record for given pk id
func (service *ArticleService) GetById(id uint64) (*Article, error) {
	var article Article
	if err := service.dataTable.FindOne(dbclient.Condition{"id": id}, &article); err != nil {
		return nil, err
	}
	return &article, nil
}

// Create creates a new record on the datastore with given struct
func (service *ArticleService) Create(a *Article) error {
	if err := service.dataTable.InsertReturning(a); err != nil {
		return err
	}
	return nil
}

// Update updates given record on the datastore by finding it with its pk
func (service *ArticleService) Update(a *Article) error {
	a.UpdatedAt = time.Now().UTC()
	if err := service.dataTable.UpdateReturning(a); err != nil {
		return err
	}
	return nil
}

// Delete deletes the given struct from database by finding it with its pk
func (service *ArticleService) Delete(a *Article) error {
	if err := service.dataTable.Delete(dbclient.Condition{"id": a.ID}); err != nil {
		return err
	}
	return nil
}
