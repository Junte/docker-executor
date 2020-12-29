build:
	@go build -o bin/executor executor/cmd/server

test:
	@go test -v ./...