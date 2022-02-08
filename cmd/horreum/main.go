package main

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	// Register stream server for event messages
	// Ideally this should live in its own package
	// with proper error handler under the cmd/ folder
	// Just left here for the demo purposes
	streamService := streamer.NewStreamer()
	go func() {
		if err = streamService.Router.Run(ctx); err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
	}()

	// Register http server and run
	srv := server.New(config, &db, streamService)
	go func() {
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
