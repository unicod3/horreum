package warehouse

import (
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/upper/db/v4"
	"time"
)

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

type Service struct {
	dataStore *dbclient.DataStore
}

func (service *Service) GetById(id uint64) (*Warehouse, error) {
	var warehouse Warehouse
	if err := service.dataStore.Find(db.Cond{"id": id}).One(&warehouse); err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (service *Service) GetAll() ([]Warehouse, error) {
	var warehouses []Warehouse
	if err := service.dataStore.Find().All(&warehouses); err != nil {
		return nil, err
	}
	return warehouses, nil
}

func (service *Service) Create(w *Warehouse) error {
	if err := service.dataStore.InsertReturning(w); err != nil {
		return err
	}
	return nil
}

func (service *Service) Update(w *Warehouse) error {
	w.UpdatedAt = time.Now().UTC()
	if err := service.dataStore.UpdateReturning(w); err != nil {
		return err
	}
	return nil
}

func (service *Service) Delete(w *Warehouse) error {
	if err := service.dataStore.Find(db.Cond{"id": w.ID}).Delete(); err != nil {
		return err
	}
	return nil
}
