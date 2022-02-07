package product

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/dbclient/mocks"
	"testing"
)

func TestProductServiceImplementsProductRepositoryInterface(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*ProductRepository)(nil), new(ProductService))
}

func TestProductService_GetAll(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	productService := &ProductService{
		dataTable: &dataTable,
	}

	products := Products{
		Product{ID: 1, SKU: "test", Price: 1025},
	}

	var w Products
	dataTable.On("FindAll", &w).Run(func(args mock.Arguments) {
		w = products
	}).Return(nil).Once()
	var productArticles []ProductArticle
	dataTable.On("LoadMany2Many", "pa.product_id as product_id, a.*",
		"product_articles pa",
		"articles a",
		"a.id = pa.article_id",
		dbclient.Condition{"pa.product_id IN ": w.IDList()}, &productArticles).Return(nil).Once()
	_, err := productService.GetAll()
	assert.Nil(err)
	assert.Equal(products, w)
}

func TestProductService_GetById(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	productService := &ProductService{
		dataTable: &dataTable,
	}

	product := Product{ID: 1, SKU: "test", Price: 1025}

	var w Product
	dataTable.On("FindOne", dbclient.Condition{"id": product.ID}, &w).Run(func(args mock.Arguments) {
		w = product
	}).Return(nil).Once()
	var productArticles []Article
	dataTable.On("LoadMany2Many", "a.*",
		"product_articles pa",
		"articles a",
		"a.id = pa.article_id",
		dbclient.Condition{"pa.product_id": w.ID},
		&productArticles).Return(nil).Once()
	_, err := productService.GetById(product.ID)
	assert.Nil(err)
	assert.Equal(product, w)
}

func TestProductService_Create(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	productService := &ProductService{
		dataTable: &dataTable,
	}

	product := Product{ID: 1, SKU: "test"}

	var w Product
	dataTable.On("InsertReturning", &product).Run(func(args mock.Arguments) {
		w = product
	}).Return(nil).Once()
	err := productService.Create(&product)
	assert.Nil(err)
	assert.Equal(product, w)
}

func TestProductService_Update(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	productService := &ProductService{
		dataTable: &dataTable,
	}

	product := Product{ID: 1, SKU: "test"}

	var w Product
	dataTable.On("UpdateReturning", &product).Run(func(args mock.Arguments) {
		w = product
	}).Return(nil).Once()
	err := productService.Update(&product)
	assert.Nil(err)
	assert.Equal(product, w)
}

func TestProductService_Delete(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	productService := &ProductService{
		dataTable: &dataTable,
	}

	product := Product{ID: 1, SKU: "test"}
	dataTable.On("Delete", dbclient.Condition{"id": product.ID}).Return(nil).Once()
	err := productService.Delete(&product)
	assert.Nil(err)
}
