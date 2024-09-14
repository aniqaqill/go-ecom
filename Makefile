build:
	@go build -o bin/go-ecom cmd/main.go

run:
	@go run cmd/main.go

test:
	@go test -v ./...

.PHONY: clean
clean:
	kill -9 $$(lsof -t -i :8080)