package warehouse

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the package's routes to the gin router
func (h *Handler) RegisterRoutes(routerGroup *gin.RouterGroup) {
	warehouses := routerGroup.Group("warehouses")
	{
		warehouses.GET("/", h.List)
		warehouses.GET("/:id", h.Get)
		warehouses.POST("/", h.Create)
		warehouses.PUT("/:id", h.Update)
		warehouses.DELETE("/:id", h.Delete)
	}
}