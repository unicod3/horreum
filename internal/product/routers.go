package product

import "github.com/gin-gonic/gin"

// RegisterHTTPRoutes registers the package's routes to the gin router
func (service *ProductService) RegisterHTTPRoutes(routerGroup *gin.RouterGroup) {
	products := routerGroup.Group("products")
	{
		products.GET("/", service.ListProducts)
		products.GET("/:id", service.GetProduct)
		products.POST("/", service.CreateProduct)
		products.PUT("/:id", service.UpdateProduct)
		products.DELETE("/:id", service.DeleteProduct)
	}
}

// RegisterHTTPRoutes registers the package's routes to the gin router
func (service *ArticleService) RegisterHTTPRoutes(routerGroup *gin.RouterGroup) {
	articles := routerGroup.Group("articles")
	{
		articles.GET("/", service.ListArticles)
		articles.GET("/:id", service.GetArticle)
		articles.POST("/", service.CreateArticle)
		articles.PUT("/:id", service.UpdateArticle)
		articles.DELETE("/:id", service.DeleteArticle)
	}
}
