package warehouse

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/dbclient/mocks"
	"testing"
)

func TestWarehouseServiceImplementsWarehouseRepositoryInterface(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*WarehouseRepository)(nil), new(WarehouseService))
}

func TestWarehouseService_GetAll(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	warehouseService := &WarehouseService{
		DataTable: &dataTable,
	}

	warehouses := []Warehouse{
		Warehouse{ID: 1, Name: "test"},
	}

	var w []Warehouse
	dataTable.On("FindAll", &w).Run(func(args mock.Arguments) {
		w = warehouses
	}).Return(nil).Once()
	_, err := warehouseService.GetAll()
	assert.Nil(err)
	assert.Equal(warehouses, w)
}

func TestWarehouseService_GetById(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	warehouseService := &WarehouseService{
		DataTable: &dataTable,
	}

	warehouse := Warehouse{ID: 1, Name: "test"}

	var w Warehouse
	dataTable.On("FindOne", dbclient.Condition{"id": warehouse.ID}, &w).Run(func(args mock.Arguments) {
		w = warehouse
	}).Return(nil).Once()
	_, err := warehouseService.GetById(warehouse.ID)
	assert.Nil(err)
	assert.Equal(warehouse, w)
}

func TestWarehouseService_Create(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	warehouseService := &WarehouseService{
		DataTable: &dataTable,
	}

	warehouse := Warehouse{ID: 1, Name: "test"}

	var w Warehouse
	dataTable.On("InsertReturning", &warehouse).Run(func(args mock.Arguments) {
		w = warehouse
	}).Return(nil).Once()
	err := warehouseService.Create(&warehouse)
	assert.Nil(err)
	assert.Equal(warehouse, w)
}

func TestWarehouseService_Update(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	warehouseService := &WarehouseService{
		DataTable: &dataTable,
	}

	warehouse := Warehouse{ID: 1, Name: "test"}

	var w Warehouse
	dataTable.On("UpdateReturning", &warehouse).Run(func(args mock.Arguments) {
		w = warehouse
	}).Return(nil).Once()
	err := warehouseService.Update(&warehouse)
	assert.Nil(err)
	assert.Equal(warehouse, w)
}

func TestWarehouseService_Delete(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	warehouseService := &WarehouseService{
		DataTable: &dataTable,
	}

	warehouse := Warehouse{ID: 1, Name: "test"}
	dataTable.On("Delete", dbclient.Condition{"id": warehouse.ID}).Return(nil).Once()
	err := warehouseService.Delete(&warehouse)
	assert.Nil(err)
}
