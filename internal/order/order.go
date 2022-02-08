package order

import (
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/streamer"
	"time"
)

const (
	OrderCreated string = "OrderCreated"
	OrderUpdated        = "OrderUpdated"
	OrderDeleted        = "OrderDeleted"
)

// OrderRepository serves as a contract over OrderService
type OrderRepository interface {
	GetAll() ([]Order, error)
	GetById(id uint64) (*Order, error)
	Create(o *Order) error
	Update(o *Order) error
	Delete(o *Order) error
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
	ProductID uint64    `json:"product_id" db:"product_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
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
		ProductID uint64 `json:"product_id"`
		Quantity  uint64 `json:"quantity"`
		UnitCost  uint64 `json:"unit_cost"`
	} `json:"lines"`
}

func (o *Order) populateLines(dataTable dbclient.DataTable) error {
	return dataTable.FindRelated("order_lines", dbclient.Condition{"order_id": o.ID}, &o.Lines)
}

func (o *Order) createLines(dataTable dbclient.DataTable) error {
	for _, line := range o.Lines {
		line.OrderID = o.ID
		err := dataTable.CreateRelated("order_lines", &line)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *Order) deleteLines(dataTable dbclient.DataTable) error {
	return dataTable.DeleteRelated("order_lines", dbclient.Condition{"order_id": o.ID})
}

// OrderService holds information about the datatable
// and implements OrderService
type OrderService struct {
	DataTable     dbclient.DataTable
	StreamChannel streamer.Channel
	StreamTopic   string
}

// GetAll returns all the records
func (service *OrderService) GetAll() ([]Order, error) {
	var orders []Order

	if err := service.DataTable.FindAll(&orders); err != nil {
		return nil, err
	}
	for i, order := range orders {
		order.populateLines(service.DataTable)
		orders[i] = order
	}

	return orders, nil
}

// GetById returns single record for given pk id
func (service *OrderService) GetById(id uint64) (*Order, error) {
	var order Order
	if err := service.DataTable.FindOne(dbclient.Condition{"id": id}, &order); err != nil {
		return nil, err
	}
	order.populateLines(service.DataTable)
	return &order, nil
}

// Create creates a new record on the datastore with given struct
func (service *OrderService) Create(o *Order) error {
	if err := service.DataTable.InsertReturning(o); err != nil {
		return err
	}
	err := o.createLines(service.DataTable)
	if err != nil {
		return err
	}

	// Publish an event on the channel
	err = service.PublishEvent(OrderCreated, o)
	if err != nil {
		return err
	}
	return nil
}

// Update updates given record on the datastore by finding it with its pk
func (service *OrderService) Update(o *Order) error {
	o.UpdatedAt = time.Now().UTC()
	if err := service.DataTable.UpdateReturning(o); err != nil {
		return err
	}
	err := o.deleteLines(service.DataTable)
	if err != nil {
		return err
	}

	err = o.createLines(service.DataTable)
	if err != nil {
		return err
	}

	// Publish an event on the channel
	err = service.PublishEvent(OrderUpdated, o)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes the given struct from database by finding it with its pk
func (service *OrderService) Delete(o *Order) error {
	if err := service.DataTable.Delete(dbclient.Condition{"id": o.ID}); err != nil {
		return err
	}

	// Publish an event on the channel
	err := service.PublishEvent(OrderDeleted, o)
	if err != nil {
		return err
	}

	return nil
}

func (service *OrderService) PublishEvent(event string, order *Order) error {
	msg, err := streamer.NewMessage(&streamer.Message{
		EventName: event,
		Data:      order,
	})
	if err != nil {
		return err
	}
	streamer.PublishMessage(service.StreamChannel, service.StreamTopic, msg)
	return nil
}
