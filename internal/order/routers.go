package order

import (
	"github.com/gin-gonic/gin"
	"github.com/unicod3/horreum/pkg/streamer"
)

// RegisterHTTPRoutes registers the package's routes to the gin router
func (service *OrderService) RegisterHTTPRoutes(routerGroup *gin.RouterGroup) {
	orders := routerGroup.Group("orders")
	{
		orders.GET("/", service.ListOrders)
		orders.GET("/:id", service.GetOrder)
		orders.POST("/", service.CreateOrder)
		orders.PUT("/:id", service.UpdateOrder)
		orders.DELETE("/:id", service.DeleteOrder)
	}
}

// RegisterEventHandlers registers the package's events handlers to streamer package
func (service *OrderService) RegisterEventHandlers(s *streamer.Stream) {
	s.RegisterHandler(
		service.StreamChannel,
		"orders",
		service.printMessages,
	)
}
