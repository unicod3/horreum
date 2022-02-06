package warehouse

import (
	"github.com/unicod3/horreum/pkg/dbclient"
	"time"
)

// Handler holds services that are exposed
type Handler struct {
	WarehouseService WarehouseRepository
}

// NewHandler returns a new Handler
func NewHandler(client *dbclient.DataStorage) *Handler {
	return &Handler{
		WarehouseService: &WarehouseService{
			(*client).NewDataCollection("warehouses"),
		},
	}
}

// Warehouse represents a record from warehouses table
type Warehouse struct {
	ID        uint64    `json:"id" uri:"id" db:"id,omitempty"`
	CreatedAt time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at,omitempty"`
	Name      string    `json:"name" db:"name"`
}

// ErrorResponse contains information about error
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// RequestBody represents the data type that needs to be sent over request
type RequestBody struct {
	Name string
}

// WarehouseRepository serves as a contract over WarehouseService
type WarehouseRepository interface {
	GetAll() ([]Warehouse, error)
	GetById(id uint64) (*Warehouse, error)
	Create(w *Warehouse) error
	Update(w *Warehouse) error
	Delete(w *Warehouse) error
}

// WarehouseService holds information about the datatable
// and implements WarehouseRepository
type WarehouseService struct {
	dataTable dbclient.DataTable
}

// GetAll returns all the records
func (service *WarehouseService) GetAll() ([]Warehouse, error) {
	var warehouses []Warehouse
	if err := service.dataTable.FindAll(&warehouses); err != nil {
		return nil, err
	}
	return warehouses, nil
}

// GetById returns single record for given pk id
func (service *WarehouseService) GetById(id uint64) (*Warehouse, error) {
	var warehouse Warehouse
	if err := service.dataTable.FindOne(dbclient.Condition{"id": id}, &warehouse); err != nil {
		return nil, err
	}
	return &warehouse, nil
}

// Create creates a new record on the datastore with given struct
func (service *WarehouseService) Create(w *Warehouse) error {
	if err := service.dataTable.InsertReturning(w); err != nil {
		return err
	}
	return nil
}

// Update updates given record on the datastore by finding it with its pk
func (service *WarehouseService) Update(w *Warehouse) error {
	w.UpdatedAt = time.Now().UTC()
	if err := service.dataTable.UpdateReturning(w); err != nil {
		return err
	}
	return nil
}

// Delete deletes the given struct from database by finding it with its pk
func (service *WarehouseService) Delete(w *Warehouse) error {
	if err := service.dataTable.Delete(dbclient.Condition{"id": w.ID}); err != nil {
		return err
	}
	return nil
}
