build:
	@go build -o bin/server cmd/main.go

test: 
	@go test -v ./...

run: build
	@./bin/server