package order

import "github.com/gin-gonic/gin"

// RegisterRoutes registers the package's routes to the gin router
func (h *Handler) RegisterRoutes(routerGroup *gin.RouterGroup) {
	orders := routerGroup.Group("orders")
	{
		orders.GET("/", h.List)
		orders.GET("/:id", h.Get)
		orders.POST("/", h.Create)
		orders.PUT("/:id", h.Update)
		orders.DELETE("/:id", h.Delete)
	}
}
