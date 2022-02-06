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
	Addr string
}

// Server contains server details
type Server struct {
	cfg *Config
	DB  *dbclient.Client
}

// New returns a new instance of the server
func New(cfg *Config, db *dbclient.Client) *Server {
	return &Server{
		cfg: cfg,
		DB:  db,
	}
}

func (srv *Server) Serve() {

	docs.SwaggerInfo_swagger.Title = "Horreum"
	docs.SwaggerInfo_swagger.Description = "Horreum, is an application to manage products and their stock information."
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	docs.SwaggerInfo_swagger.Host = "localhost:8080"

	warehouseHandler := warehouse.NewHandler(srv.DB)
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		warehouseHandler.RegisterRoutes(v1)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(srv.cfg.Addr)
}
