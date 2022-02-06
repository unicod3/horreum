package product

import "github.com/gin-gonic/gin"

// RegisterRoutes registers the package's routes to the gin router
func (h *Handler) RegisterRoutes(routerGroup *gin.RouterGroup) {
	products := routerGroup.Group("products")
	{
		products.GET("/", h.ListProduct)
		products.GET("/:id", h.GetProduct)
		products.POST("/", h.CreateProduct)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)
	}
	articles := routerGroup.Group("articles")
	{
		articles.GET("/", h.ListArticle)
		articles.GET("/:id", h.GetArticle)
		articles.POST("/", h.CreateArticle)
		articles.PUT("/:id", h.UpdateArticle)
		articles.DELETE("/:id", h.DeleteArticle)
	}
}
