package main

import (
	"github.com/unicod3/horreum/api/server"
)

func main() {
	srv := server.New(&server.Config{
		Addr: ":8080",
	})

	srv.Serve()
}
