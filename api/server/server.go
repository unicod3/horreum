package server

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/unicod3/horreum/api/docs"
	"github.com/unicod3/horreum/internal/warehouse"
	"github.com/unicod3/horreum/pkg/dbclient"
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
	cfg       *Config
	DataStore *dbclient.DataStorage
}

// New returns a new instance of the server
func New(cfg *Config, db *dbclient.DataStorage) *Server {
	return &Server{
		cfg:       cfg,
		DataStore: db,
	}
}

var ginRouter = gin.Default()

func (srv *Server) Serve() {
	docs.SwaggerInfo_swagger.Title = srv.cfg.SwaggerTitle
	docs.SwaggerInfo_swagger.Description = srv.cfg.SwaggerDescription
	docs.SwaggerInfo_swagger.BasePath = srv.cfg.BasePath
	docs.SwaggerInfo_swagger.Host = srv.cfg.Addr

	router := registerGinRouter(srv.cfg.BasePath)

	warehouseHandler := warehouse.NewHandler(srv.DataStore)
	warehouseHandler.RegisterRoutes(router)

	ginRouter.Run(srv.cfg.Addr)
}

func registerGinRouter(basePath string) *gin.RouterGroup {
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return ginRouter.Group(basePath)
}
