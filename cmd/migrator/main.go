package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	_ "github.com/unicod3/horreum/pkg/dbclient/migrations"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}
	command := args[0]

	dbString := os.Getenv("MIGRATOR_CONN")
	dbDriver := os.Getenv("DATABASE_DRIVER")

	db, err := goose.OpenDBWithDriver(dbDriver, dbString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, ".", arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
