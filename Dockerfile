FROM golang:1.17-alpine as development
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Expose port
EXPOSE 8080

# Start app
CMD go run cmd/migrator/main.go up && go run cmd/horreum/main.go
