package order

import (
	"github.com/gin-gonic/gin"
	"github.com/unicod3/horreum/pkg/streamer"
)

// RegisterHTTPRoutes registers the package's routes to the gin router
func (h *Handler) RegisterHTTPRoutes(routerGroup *gin.RouterGroup) {
	orders := routerGroup.Group("orders")
	{
		orders.GET("/", h.List)
		orders.GET("/:id", h.Get)
		orders.POST("/", h.Create)
		orders.PUT("/:id", h.Update)
		orders.DELETE("/:id", h.Delete)
	}
}

// RegisterEventHandlers registers the package's events handlers to streamer package
func (h *Handler) RegisterEventHandlers(s *streamer.Stream) {
	s.RegisterHandler(
		h.OrderService.StreamChannel,
		"orders",
		h.printMessages,
	)
}
