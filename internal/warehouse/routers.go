package warehouse

import (
	"github.com/gin-gonic/gin"
)

// RegisterHTTPRoutes registers the package's routes to the gin router
func (service WarehouseService) RegisterHTTPRoutes(routerGroup *gin.RouterGroup) {
	warehouses := routerGroup.Group("warehouses")
	{
		warehouses.GET("/", service.ListWarehouses)
		warehouses.GET("/:id", service.GetWarehouse)
		warehouses.POST("/", service.CreateWarehouse)
		warehouses.PUT("/:id", service.UpdateWarehouse)
		warehouses.DELETE("/:id", service.DeleteWarehouse)
	}
}
