build:
	@mkdir -p bin
	@go build -o bin/ecom cmd/main.go

test: 
	@go test -v ./...

run: build
	@./bin/ecom