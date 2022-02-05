package server

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/unicod3/horreum/api/docs"
)

// Config provides the configuration for the API server
type Config struct {
	Addr string
}

// Server contains server details
type Server struct {
	cfg *Config
}

// New returns a new instance of the server
func New(cfg *Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Serve() {
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	docs.SwaggerInfo_swagger.Title = "Horreum"
	docs.SwaggerInfo_swagger.Description = "Horreum, is an application to manage products and their stock information."

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(s.cfg.Addr)
}
