run:
	go run cmd/horreum/main.go

tidy:
	go mod tidy


# OPEN API swagger generator
# This command will generate docs under the api/docs folder
swag:
	swag init --dir=./api/server/ --generalInfo=server.go --output=./api/docs
