package warehouse

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	warehouses := routerGroup.Group("warehouses")
	{
		warehouses.GET("/", List)
		warehouses.GET("/:id", Get)
		warehouses.POST("/", Create)
		warehouses.PUT("/:id", Update)
		warehouses.DELETE("/:id", Delete)
	}
}
