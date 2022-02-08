package order

import (
	"github.com/gin-gonic/gin"
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
