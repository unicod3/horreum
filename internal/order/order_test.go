package order

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/dbclient/mocks"
	"testing"
)

func TestOrderServiceImplementsOrderRepositoryInterface(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*OrderRepository)(nil), new(OrderService))
}

func TestOrderService_GetAll(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	orderService := &OrderService{
		DataTable: &dataTable,
	}

	orders := []Order{
		Order{ID: 1, Customer: "test"},
		Order{ID: 2, Customer: "test 2"},
	}

	var w []Order
	dataTable.On("FindAll", &w).Run(func(args mock.Arguments) {
		w = orders
	}).Return(nil).Once()
	_, err := orderService.GetAll()
	assert.Nil(err)
	assert.Equal(orders, w)
}

func TestOrderService_GetById(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	orderService := &OrderService{
		DataTable: &dataTable,
	}

	order := Order{ID: 1, Customer: "test"}

	var w Order
	dataTable.On("FindOne", dbclient.Condition{"id": order.ID}, &w).Run(func(args mock.Arguments) {
		w = order
	}).Return(nil).Once()
	dataTable.On("FindRelated", "order_lines", dbclient.Condition{"order_id": w.ID}, &w.Lines).
		Run(func(args mock.Arguments) {
			w = order
		}).Return(nil).Once()

	_, err := orderService.GetById(order.ID)
	assert.Nil(err)
	assert.Equal(order, w)
}

func TestOrderService_Create(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	orderService := &OrderService{
		DataTable: &dataTable,
	}

	order := Order{ID: 1, Customer: "test"}

	var w Order
	dataTable.On("InsertReturning", &order).Run(func(args mock.Arguments) {
		w = order
	}).Return(nil).Once()
	err := orderService.Create(&order)
	assert.Nil(err)
	assert.Equal(order, w)
}

func TestOrderService_Update(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	orderService := &OrderService{
		DataTable: &dataTable,
	}

	order := Order{ID: 1, Customer: "test"}

	var w Order
	dataTable.On("UpdateReturning", &order).Run(func(args mock.Arguments) {
		w = order
	}).Return(nil).Once()

	dataTable.On("DeleteRelated", "order_lines", dbclient.Condition{"order_id": order.ID}).Return(nil).Once()

	err := orderService.Update(&order)
	assert.Nil(err)
	assert.Equal(order, w)
}

func TestOrderService_Delete(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	orderService := &OrderService{
		DataTable: &dataTable,
	}

	order := Order{ID: 1, Customer: "test"}
	dataTable.On("Delete", dbclient.Condition{"id": order.ID}).Return(nil).Once()
	err := orderService.Delete(&order)
	assert.Nil(err)
}
