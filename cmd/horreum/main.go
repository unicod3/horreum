package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/unicod3/horreum/api/server"
	"github.com/unicod3/horreum/pkg/dbclient"
	"github.com/unicod3/horreum/pkg/streamer"
	"log"
	"os"
	"os/signal"
	"syscall"
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
		Addr:               ":8080",
		SwaggerURL:         "localhost:8080",
		BasePath:           "/api/v1",
		SwaggerTitle:       "Horreum",
		SwaggerDescription: "Horreum, is an application to manage products and their stock information.",
	}

	streamService := streamer.NewStreamer()

	// Register http server and run
	go func() {
		srv := server.New(config, &db, streamService)
		if err = srv.Serve(); err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
