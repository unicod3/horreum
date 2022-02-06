package main

import (
	"github.com/joho/godotenv"
	"github.com/unicod3/horreum/api/server"
	"github.com/unicod3/horreum/pkg/dbclient"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env file is missing")
	}

	db := dbclient.NewPostgresClient(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASS"))

	config := &server.Config{
		Addr:               "localhost:8080",
		BasePath:           "/api/v1",
		SwaggerTitle:       "Horreum",
		SwaggerDescription: "Horreum, is an application to manage products and their stock information.",
	}

	srv := server.New(config, &db)
	srv.Serve()
}
