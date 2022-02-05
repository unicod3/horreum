run:
	go run cmd/horreum/main.go

tidy:
	go mod tidy
	go mod vendor


# OPEN API swagger generator
# This command will generate docs under the api/docs folder
swag:
	swag init --dir=./api/server/ --generalInfo=server.go --output=./api/docs


# Goose database migration tool commands
migrate-up:
	go run cmd/migrator/main.go up

migrate-status:
	go run cmd/migrator/main.go status

migrate-down:
	go run cmd/migrator/main.go down

migrate-version:
	go run cmd/migrator/main.go version