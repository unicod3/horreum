package server

import (
	"github.com/unicod3/horreum/internal/article"
	"github.com/unicod3/horreum/internal/order"
	"github.com/unicod3/horreum/internal/product"
	"github.com/unicod3/horreum/internal/warehouse"
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/streamer"
)

// Handler holds services that are exposed
type Handler struct {
	WarehouseService *warehouse.WarehouseService
	OrderService     *order.OrderService
	ArticleService   *article.ArticleService
	ProductService   *product.ProductService
}

// NewHandler returns a new Handler
func NewHandler(client *dbclient.DataStorage, streamChannel streamer.Channel) *Handler {
	return &Handler{
		OrderService: &order.OrderService{
			DataTable:     (*client).NewDataCollection("orders"),
			StreamChannel: streamChannel,
			StreamTopic:   "orders",
		},
		WarehouseService: &warehouse.WarehouseService{
			DataTable:     (*client).NewDataCollection("warehouses"),
			StreamChannel: streamChannel,
			StreamTopic:   "warehouses",
		},
		ArticleService: &article.ArticleService{
			DataTable:     (*client).NewDataCollection("articles"),
			StreamChannel: streamChannel,
			StreamTopic:   "articles",
		},
		ProductService: &product.ProductService{
			DataTable:     (*client).NewDataCollection("products"),
			StreamChannel: streamChannel,
			StreamTopic:   "products",
		},
	}
}
