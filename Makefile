APP_NAME=101-golang

start:
	go run main.go $(mod)

build:
	go build -o bin/$(APP_NAME) main.go

clean:
	rm -rf bin

help:
	@echo "Usage:"
	@echo "  make start mod=rate_limit   # Run rate limiter"
	@echo "  make build                  # Build binary"
	@echo "  make clean                  # Clean build artifacts"
