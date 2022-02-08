package product

import (
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/streamer"
	"sort"
	"time"
)

// ProductRepository serves as a contract over ArticleService
type ProductRepository interface {
	GetAll() (Products, error)
	GetById(uint64) (*Product, error)
	Create(*Product) (*Product, error)
	Update(*Product) (*Product, error)
	Delete(*Product) error
}

// Product represents a record from products table
type Product struct {
	ID                uint64    `json:"id" uri:"id" db:"id,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	Name              string    `json:"name" db:"name"`
	Price             int64     `json:"price" db:"price"`
	SellableInventory int64     `json:"sellable_inventory,omitempty" db:"-"`
	Articles          []Article `json:"articles" db:"-"`
}

func (p *Product) CalculateSellableInventory() {
	if len(p.Articles) == 0 {
		p.SellableInventory = 0
		return
	}
	articles := p.Articles
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].AvailableInventory < articles[j].AvailableInventory
	})
	minInventoryOfArticles := articles[0].AvailableInventory
	p.SellableInventory = minInventoryOfArticles
}

func (p *Product) IncreaseStockBy(articleService *ArticleService, quantity int64) error {
	if quantity == 0 {
		return nil
	}

	for _, article := range p.Articles {
		decreaseAmount := article.AmountOf * quantity
		newStock := article.Stock + decreaseAmount
		err := articleService.Update(&Article{
			ID:    article.ID,
			Stock: newStock,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Product) DecreaseStockBy(articleService *ArticleService, quantity int64) error {
	if quantity == 0 {
		return nil
	}

	for _, article := range p.Articles {
		decreaseAmount := article.AmountOf * quantity
		newStock := article.Stock - decreaseAmount
		err := articleService.Update(&Article{
			ID:    article.ID,
			Stock: newStock,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type ProductArticleRelation struct {
	ProductID uint64 `db:"product_id"`
	ArticleID uint64 `db:"article_id"`
	AmountOf  int64  `db:"amount_of"`
}

type ProductArticle struct {
	ProductID uint64 `db:"product_id"`
	Article   `db:",inline"`
}

// Products holds multiple Product
type Products []Product

// ProductMap serves a map to m2m relation
type ProductMap map[uint64]Product

// IDList gets slice of ids from a Products struct
func (products Products) IDList() []uint64 {
	var list []uint64
	for _, product := range products {
		list = append(list, product.ID)
	}
	return list
}

// ConvertToMap builds up a ProductMap object
func (products Products) ConvertToMap() ProductMap {
	m := make(ProductMap, len(products))
	for _, product := range products {
		m[product.ID] = product
	}
	return m
}

// Products returns a new Products struct from ProductMap
func (productMap ProductMap) Products() Products {
	var products Products
	for _, product := range productMap {
		products = append(products, product)
	}
	return products
}

// ErrorResponse contains information about error
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ProductRequestBody represents the data type that needs to be sent over request
type ProductRequestBody struct {
	Name     string `json:"name"`
	Price    uint64 `json:"price"`
	Articles []struct {
		ID        uint64 `json:"id"`
		ProductID uint64 `json:"-"`
		AmountOf  int64  `json:"amount_of"`
	} `json:"articles"`
}

// ProductService holds information about the datatable
// and implements ArticleService
type ProductService struct {
	DataTable     dbclient.DataTable
	StreamChannel streamer.Channel
	StreamTopic   string
}

// GetAll returns all the records
func (service *ProductService) GetAll() (Products, error) {
	var products Products
	err := service.DataTable.FindAll(&products)
	if err != nil {
		return nil, err
	}
	products, err = service.populateArticles(products)
	if err != nil {
		return nil, err
	}
	products, err = service.populateSellableInventory(products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetById returns single record for given pk id
func (service *ProductService) GetById(id uint64) (*Product, error) {
	var product Product
	if err := service.DataTable.FindOne(dbclient.Condition{"id": id}, &product); err != nil {
		return nil, err
	}
	err := service.populateArticle(&product)
	if err != nil {
		return nil, err
	}
	(&product).CalculateSellableInventory()
	return &product, nil
}

// Create creates a new record on the datastore with given struct
func (service *ProductService) Create(p *Product) (*Product, error) {
	if err := service.DataTable.InsertReturning(p); err != nil {
		return nil, err
	}
	err := service.syncArticles(p)
	if err != nil {
		return nil, err
	}
	p, _ = service.GetById(p.ID)
	return p, nil
}

// Update updates given record on the datastore by finding it with its pk
func (service *ProductService) Update(p *Product) (*Product, error) {
	p.UpdatedAt = time.Now().UTC()
	if err := service.DataTable.UpdateReturning(p); err != nil {
		return nil, err
	}
	err := service.syncArticles(p)
	if err != nil {
		return nil, err
	}
	p, _ = service.GetById(p.ID)
	return p, nil
}

// Delete deletes the given struct from database by finding it with its pk
func (service *ProductService) Delete(p *Product) error {
	if err := service.DataTable.Delete(dbclient.Condition{"id": p.ID}); err != nil {
		return err
	}
	return nil
}

func (service *ProductService) populateArticle(product *Product) error {
	var productArticles []Article
	err := service.DataTable.LoadMany2Many(
		"a.*, pa.amount_of as amount_of",
		"product_articles pa",
		"articles a",
		"a.id = pa.article_id",
		dbclient.Condition{"pa.product_id": product.ID},
		&productArticles)

	if err != nil {
		return err
	}
	for i, article := range productArticles {
		(&article).CalculateAvailableInventory()
		productArticles[i] = article
	}

	product.Articles = productArticles
	return nil
}

func (service *ProductService) populateArticles(products Products) (Products, error) {
	var productArticles []ProductArticle
	err := service.DataTable.LoadMany2Many(
		"pa.product_id as product_id, a.*, pa.amount_of as amount_of",
		"product_articles pa",
		"articles a",
		"a.id = pa.article_id",
		dbclient.Condition{"pa.product_id IN ": products.IDList()},
		&productArticles)

	if err != nil {
		return nil, err
	}
	var productMap = products.ConvertToMap()
	for _, productArticle := range productArticles {
		product := productMap[productArticle.ProductID]

		article := Article{
			ID:        productArticle.ID,
			CreatedAt: productArticle.CreatedAt,
			UpdatedAt: productArticle.UpdatedAt,
			Name:      productArticle.Name,
			Stock:     productArticle.Stock,
			AmountOf:  productArticle.AmountOf,
		}
		article.CalculateAvailableInventory()
		product.Articles = append(product.Articles, article)
		productMap[productArticle.ProductID] = product
	}

	return productMap.Products(), nil
}

func (service *ProductService) syncArticles(p *Product) error {
	err := service.DataTable.DeleteRelated("product_articles", dbclient.Condition{"product_id": p.ID})
	if err != nil {
		return err
	}
	for _, article := range p.Articles {
		err = service.DataTable.CreateRelated("product_articles", &ProductArticleRelation{
			ProductID: p.ID,
			ArticleID: article.ID,
			AmountOf:  article.AmountOf,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *ProductService) populateSellableInventory(products Products) (Products, error) {
	for i, product := range products {
		(&product).CalculateSellableInventory()
		products[i] = product
	}
	return products, nil
}
