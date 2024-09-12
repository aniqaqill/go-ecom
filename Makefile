build:
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

run:
	@go run cmd/$(APP_NAME)/main.go

test:
	@go test -v ./...

.PHONY: clean
clean:
    kill -9 $(lsof -t -i :8080)