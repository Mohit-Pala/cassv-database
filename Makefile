BINARY_NAME=cassv

all: run

build:
	@go build -o $(BINARY_NAME) .

run: build
	@./$(BINARY_NAME)

clean:
	@rm -f $(BINARY_NAME)
	@rm -rf cassv_filesystem

test:
	@go test ./...

.PHONY: all build run clean test
