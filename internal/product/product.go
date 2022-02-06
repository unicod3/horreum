package product

import (
	"github.com/unicod3/horreum/pkg/dbclient"
	"time"
)

// Handler holds services that are exposed
type Handler struct {
	ArticleService ArticleRepository
	ProductService ProductRepository
}

// NewHandler returns a new Handler
func NewHandler(client *dbclient.DataStorage) *Handler {
	return &Handler{
		ArticleService: &ArticleService{
			dataTable: (*client).NewDataCollection("articles"),
		},
		ProductService: &ProductService{
			dataTable: (*client).NewDataCollection("products"),
		},
	}
}

// ProductRepository serves as a contract over ArticleService
type ProductRepository interface {
	GetAll() ([]Product, error)
	GetById(uint64) (*Product, error)
	Create(*Product) error
	Update(*Product) error
	Delete(*Product) error
}

// Product represents a record from products table
type Product struct {
	ID        uint64    `json:"id" uri:"id" db:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	SKU       string    `json:"sku" db:"sku"`
	Price     int64     `json:"price" db:"price"`
}

// ErrorResponse contains information about error
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ProductRequestBody represents the data type that needs to be sent over request
type ProductRequestBody struct {
	SKU   string `json:"sku"`
	Price uint64 `json:"price"`
}

// ProductService holds information about the datatable
// and implements ArticleService
type ProductService struct {
	dataTable dbclient.DataTable
}

// GetAll returns all the records
func (service *ProductService) GetAll() ([]Product, error) {
	var products []Product

	if err := service.dataTable.FindAll(&products); err != nil {
		return nil, err
	}
	return products, nil
}

// GetById returns single record for given pk id
func (service *ProductService) GetById(id uint64) (*Product, error) {
	var product Product
	if err := service.dataTable.FindOne(dbclient.Condition{"id": id}, &product); err != nil {
		return nil, err
	}
	return &product, nil
}

// Create creates a new record on the datastore with given struct
func (service *ProductService) Create(p *Product) error {
	if err := service.dataTable.InsertReturning(p); err != nil {
		return err
	}
	return nil
}

// Update updates given record on the datastore by finding it with its pk
func (service *ProductService) Update(p *Product) error {
	p.UpdatedAt = time.Now().UTC()
	if err := service.dataTable.UpdateReturning(p); err != nil {
		return err
	}
	return nil
}

// Delete deletes the given struct from database by finding it with its pk
func (service *ProductService) Delete(p *Product) error {
	if err := service.dataTable.Delete(dbclient.Condition{"id": p.ID}); err != nil {
		return err
	}
	return nil
}
