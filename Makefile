build:
	@go build -o bin/asana-clone-be

run: build
	@./bin/asana-clone-be

test:
	@go test -v ./...