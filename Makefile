build:
	@go build -o bin/go-ecom cmd/main.go

run:
	@go run cmd/main.go

test:
	@go test -v ./...

.PHONY: clean
clean:
	kill -9 $$(lsof -t -i :8080)

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down