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
	GetAll() (Products, error)
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
	Articles  []Article `json:"articles" db:"-"`
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

func (service *ProductService) populateArticle(product *Product) error {
	var productArticles []Article
	err := service.dataTable.LoadMany2Many(
		"a.*",
		"product_articles pa",
		"articles a",
		"a.id = pa.article_id",
		dbclient.Condition{"pa.product_id": product.ID},
		&productArticles)

	if err != nil {
		return err
	}
	product.Articles = productArticles
	return nil
}

func (service *ProductService) populateArticles(products Products) (Products, error) {
	var productArticles []ProductArticle
	err := service.dataTable.LoadMany2Many(
		"pa.product_id as product_id, a.*",
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
			SKU:       productArticle.SKU,
			Quantity:  productArticle.Quantity,
		}
		product.Articles = append(product.Articles, article)
		productMap[productArticle.ProductID] = product
	}

	return productMap.Products(), nil
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
	SKU   string `json:"sku"`
	Price uint64 `json:"price"`
}

// ProductService holds information about the datatable
// and implements ArticleService
type ProductService struct {
	dataTable dbclient.DataTable
}
type ProductArticle struct {
	ProductID uint64 `db:"product_id"`
	Article   `db:",inline"`
}

// GetAll returns all the records
func (service *ProductService) GetAll() (Products, error) {
	var products Products
	err := service.dataTable.FindAll(&products)
	if err != nil {
		return nil, err
	}
	products, err = service.populateArticles(products)
	if err != nil {
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
	err := service.populateArticle(&product)
	if err != nil {
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
