package order

import (
	"github.com/unicod3/horreum/pkg/dbclient"
	"time"
)

// OrderRepository serves as a contract over OrderService
type OrderRepository interface {
	GetAll() ([]Order, error)
	GetById(id uint64) (*Order, error)
	Create(o *Order) error
	Update(o *Order) error
	Delete(o *Order) error
}

// Handler holds services that are exposed
type Handler struct {
	OrderService OrderRepository
}

// NewHandler returns a new Handler
func NewHandler(client *dbclient.DataStorage) *Handler {
	return &Handler{
		OrderService: &OrderService{
			(*client).NewDataCollection("orders"),
		},
	}
}

// Order represents a record from orders table
type Order struct {
	ID          uint64      `json:"id" uri:"id" db:"id,omitempty"`
	WarehouseID uint64      `json:"warehouse_id" db:"warehouse_id,inline"`
	CreatedAt   time.Time   `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	Customer    string      `json:"customer" db:"customer"`
	Lines       []OrderLine `json:"lines" db:"-"`
}

// OrderLine represents a record from order_lines table
type OrderLine struct {
	ID        uint64    `json:"id" db:"id,omitempty"`
	OrderID   uint64    `json:"-" db:"order_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	SKU       string    `json:"sku" db:"sku"`
	Quantity  uint64    `json:"quantity" db:"quantity"`
	UnitCost  uint64    `json:"unit_cost" db:"unit_cost"`
}

// ErrorResponse contains information about error
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// RequestBody represents the data type that needs to be sent over request
type RequestBody struct {
	Customer    string `json:"customer"`
	WarehouseID uint64 `json:"warehouse_id"`
	Lines       []struct {
		SKU      string `json:"sku"`
		Quantity uint64 `json:"quantity"`
		UnitCost uint64 `json:"unit_cost"`
	} `json:"lines"`
}

func (o *Order) populateLines(dataTable dbclient.DataTable) error {
	return dataTable.Session().
		Collection("order_lines").
		Find(dbclient.Condition{"order_id": o.ID}).
		All(&o.Lines)
}

func (o *Order) createLines(dataTable dbclient.DataTable) error {
	for _, line := range o.Lines {
		line.OrderID = o.ID
		err := dataTable.Session().
			Collection("order_lines").
			InsertReturning(&line)
		if err != nil {
			return err
		}
	}
	return nil
}
func (o *Order) deleteLines(dataTable dbclient.DataTable) error {
	return dataTable.Session().
		Collection("order_lines").
		Find(dbclient.Condition{"order_id": o.ID}).Delete()
}

// OrderService holds information about the datatable
// and implements OrderService
type OrderService struct {
	dataTable dbclient.DataTable
}

// GetAll returns all the records
func (service *OrderService) GetAll() ([]Order, error) {
	var orders []Order

	if err := service.dataTable.FindAll(&orders); err != nil {
		return nil, err
	}
	for i, order := range orders {
		order.populateLines(service.dataTable)
		orders[i] = order
	}

	return orders, nil
}

// GetById returns single record for given pk id
func (service *OrderService) GetById(id uint64) (*Order, error) {
	var order Order
	if err := service.dataTable.FindOne(dbclient.Condition{"id": id}, &order); err != nil {
		return nil, err
	}
	order.populateLines(service.dataTable)
	return &order, nil
}

// Create creates a new record on the datastore with given struct
func (service *OrderService) Create(o *Order) error {
	if err := service.dataTable.InsertReturning(o); err != nil {
		return err
	}
	err := o.createLines(service.dataTable)
	if err != nil {
		return err
	}
	return nil
}

// Update updates given record on the datastore by finding it with its pk
func (service *OrderService) Update(o *Order) error {
	o.UpdatedAt = time.Now().UTC()
	if err := service.dataTable.UpdateReturning(o); err != nil {
		return err
	}
	err := o.deleteLines(service.dataTable)
	if err != nil {
		return err
	}

	err = o.createLines(service.dataTable)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes the given struct from database by finding it with its pk
func (service *OrderService) Delete(o *Order) error {
	if err := service.dataTable.Delete(dbclient.Condition{"id": o.ID}); err != nil {
		return err
	}
	return nil
}
