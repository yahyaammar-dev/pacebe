build:
	@go build -o bin/base cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/base

swagger:
	swag init -g cmd/main.go