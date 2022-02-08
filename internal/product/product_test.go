package product

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/unicod3/horreum/internal/article"
	articleMock "github.com/unicod3/horreum/internal/article/mocks"
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/dbclient/mocks"
	"github.com/upper/db/v4"
	"testing"
)

func TestProduct_CalculateSellableInventory(t *testing.T) {
	assert := assert.New(t)

	t.Run("Test can return zero for empty article", func(t *testing.T) {
		product := &Product{}
		product.CalculateSellableInventory()
		assert.Equal(int64(0), product.SellableInventory)
	})

	t.Run("Test can get the minimum article inventory", func(t *testing.T) {
		product := &Product{
			Articles: []article.Article{
				{
					AvailableInventory: -4,
				},
				{
					AvailableInventory: 5,
				},
				{
					AvailableInventory: 10,
				},
			},
		}

		product.CalculateSellableInventory()
		assert.Equal(int64(-4), product.SellableInventory)
	})
}

func TestProduct_DecreaseStockBy(t *testing.T) {
	assert := assert.New(t)

	t.Run("Test can decrease stock", func(t *testing.T) {
		art1 := article.Article{
			ID:       1,
			Stock:    10,
			AmountOf: 3,
		}
		art2 := article.Article{
			ID:       2,
			Stock:    5,
			AmountOf: 2,
		}
		art3 := article.Article{
			ID:       3,
			Stock:    10,
			AmountOf: 3,
		}

		product := &Product{
			Articles: []article.Article{
				art1, art2, art3,
			},
		}

		orderQuantity := int64(2)
		newArt1 := article.Article{
			ID:    uint64(1),
			Stock: art1.Stock - (art1.AmountOf * orderQuantity),
		}

		newArt2 := article.Article{
			ID:    uint64(2),
			Stock: art2.Stock - (art2.AmountOf * orderQuantity),
		}

		newArt3 := article.Article{
			ID:    uint64(3),
			Stock: art3.Stock - (art3.AmountOf * orderQuantity),
		}
		articleService := &articleMock.ArticleRepository{}
		articleService.On("Update", &newArt1).Return(nil)
		articleService.On("Update", &newArt2).Return(nil)
		articleService.On("Update", &newArt3).Return(nil)

		err := product.DecreaseStockBy(articleService, orderQuantity)
		if err != nil {
			return
		}
		assert.Equal(int64(0), product.SellableInventory)
	})
}

func TestProduct_IncreaseStockBy(t *testing.T) {
	assert := assert.New(t)

	t.Run("Test can increase stock", func(t *testing.T) {
		art1 := article.Article{
			ID:       1,
			Stock:    10,
			AmountOf: 3,
		}
		art2 := article.Article{
			ID:       2,
			Stock:    5,
			AmountOf: 2,
		}
		art3 := article.Article{
			ID:       3,
			Stock:    10,
			AmountOf: 3,
		}

		product := &Product{
			Articles: []article.Article{
				art1, art2, art3,
			},
		}

		orderQuantity := int64(2)
		newArt1 := article.Article{
			ID:    uint64(1),
			Stock: art1.Stock + (art1.AmountOf * orderQuantity),
		}

		newArt2 := article.Article{
			ID:    uint64(2),
			Stock: art2.Stock + (art2.AmountOf * orderQuantity),
		}

		newArt3 := article.Article{
			ID:    uint64(3),
			Stock: art3.Stock + (art3.AmountOf * orderQuantity),
		}
		articleService := &articleMock.ArticleRepository{}
		articleService.On("Update", &newArt1).Return(nil)
		articleService.On("Update", &newArt2).Return(nil)
		articleService.On("Update", &newArt3).Return(nil)

		err := product.IncreaseStockBy(articleService, orderQuantity)
		if err != nil {
			return
		}
		assert.Equal(int64(0), product.SellableInventory)
	})
}

func TestProductServiceImplementsProductRepositoryInterface(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*ProductRepository)(nil), new(ProductService))
}

func TestProductService_GetAll(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	productService := &ProductService{
		DataTable: &dataTable,
	}

	products := Products{
		Product{ID: 1, Name: "test", Price: 1025},
	}

	var w Products
	dataTable.On("FindAll", &w).Run(func(args mock.Arguments) {
		w = products
	}).Return(nil).Once()
	var productArticles []ProductArticle
	dataTable.On("LoadMany2Many", "pa.product_id as product_id, a.*, pa.amount_of as amount_of",
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
		DataTable: &dataTable,
	}

	product := Product{ID: 1, Name: "test", Price: 1025}

	var w Product
	dataTable.On("FindOne", dbclient.Condition{"id": product.ID}, &w).Run(func(args mock.Arguments) {
		w = product
	}).Return(nil).Once()
	var productArticles []article.Article
	dataTable.On("LoadMany2Many", "a.*, pa.amount_of as amount_of",
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
		DataTable: &dataTable,
	}

	productID := uint64(1)
	product := Product{
		Name:  "test",
		Price: 1000,
	}

	dataTable.On("InsertReturning", &product).
		Return(func(data interface{}) error {
			(&product).ID = productID
			return nil
		}).Once()
	dataTable.On("DeleteRelated", "product_articles", dbclient.Condition{"product_id": productID}).
		Return(nil).Once()

	result := Product{}
	dataTable.On("FindOne", dbclient.Condition{"id": productID}, &result).
		Return(func(cond db.Cond, dataAddress interface{}) error {
			(&product).ID = productID
			(&product).Name = product.Name
			(&product).Price = product.Price
			return nil
		}).Once()
	dataTable.On("LoadMany2Many",
		"a.*, pa.amount_of as amount_of",
		"product_articles pa",
		"articles a",
		"a.id = pa.article_id",
		dbclient.Condition{"pa.product_id": product.ID},
		new([]article.Article)).
		Return(nil).Once()
	p, err := productService.Create(&product)
	assert.Nil(err)
	assert.Equal(result, *p)
}

func TestProductService_Update(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	productService := &ProductService{
		DataTable: &dataTable,
	}
	productID := uint64(1)
	product := Product{
		Name:  "test",
		Price: 10,
	}

	dataTable.On("UpdateReturning", &product).Return(nil).
		Return(func(data interface{}) error {
			(&product).ID = productID
			return nil
		}).Once()
	dataTable.On("DeleteRelated", "product_articles", dbclient.Condition{"product_id": productID}).
		Return(nil).Once()

	result := Product{}
	dataTable.On("FindOne", dbclient.Condition{"id": productID}, &result).
		Return(func(cond db.Cond, dataAddress interface{}) error {
			(&product).ID = productID
			(&product).Name = product.Name
			(&product).Price = product.Price
			return nil
		}).Once()
	dataTable.On("LoadMany2Many",
		"a.*, pa.amount_of as amount_of",
		"product_articles pa",
		"articles a",
		"a.id = pa.article_id",
		dbclient.Condition{"pa.product_id": product.ID},
		new([]article.Article)).
		Return(nil).Once()
	p, err := productService.Update(&product)
	assert.Nil(err)
	assert.Equal(result, *p)
}

func TestProductService_Delete(t *testing.T) {
	assert := assert.New(t)

	dataTable := mocks.DataTable{}
	productService := &ProductService{
		DataTable: &dataTable,
	}

	product := Product{ID: 1, Name: "test"}
	dataTable.On("Delete", dbclient.Condition{"id": product.ID}).Return(nil).Once()
	err := productService.Delete(&product)
	assert.Nil(err)
}
