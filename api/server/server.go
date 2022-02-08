package server

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/unicod3/horreum/api/docs"
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/streamer"
)

// Config provides the configuration for the API server
type Config struct {
	SwaggerTitle       string
	SwaggerDescription string
	BasePath           string
	Addr               string
}

// Server contains server details
type Server struct {
	cfg           *Config
	DataStore     *dbclient.DataStorage
	StreamService *streamer.Stream
}

// New returns a new instance of the server
func New(cfg *Config, db *dbclient.DataStorage, streamService *streamer.Stream) *Server {
	return &Server{
		cfg:           cfg,
		DataStore:     db,
		StreamService: streamService,
	}
}

var ginRouter = gin.Default()

// Serve registers the ginRouter and runs it
func (srv *Server) Serve() error {
	docs.SwaggerInfo_swagger.Title = srv.cfg.SwaggerTitle
	docs.SwaggerInfo_swagger.Description = srv.cfg.SwaggerDescription
	docs.SwaggerInfo_swagger.BasePath = srv.cfg.BasePath
	docs.SwaggerInfo_swagger.Host = srv.cfg.Addr

	router := registerGinRouter(srv.cfg.BasePath)

	// Register all the internal services
	handler := NewHandler(srv.DataStore, streamer.NewChannel())
	handler.RegisterEventHandlers(srv.StreamService)

	handler.OrderService.RegisterHTTPRoutes(router)
	handler.WarehouseService.RegisterHTTPRoutes(router)
	handler.ArticleService.RegisterHTTPRoutes(router)
	handler.ProductService.RegisterHTTPRoutes(router)

	// Ideally this should live in its own package
	// with proper error handler under the cmd/ folder
	// Just left here for the demo purposes
	go srv.StreamService.Router.Run(context.Background())

	return ginRouter.Run(srv.cfg.Addr)
}

func registerGinRouter(basePath string) *gin.RouterGroup {
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return ginRouter.Group(basePath)
}
