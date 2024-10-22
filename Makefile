# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	air

run_frontend:
	cd frontend && npm run build:css


start_service_docker:
	sudo /etc/init.d/docker start

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

start_docker:
	sudo docker start postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root shortvideo

dropdb:
	docker exec -it postgres dropdb shortvideo

# Create DB container
# docker-run:
# 	@if docker compose up 2>/dev/null; then \
# 		: ; \
# 	else \
# 		echo "Falling back to Docker Compose V1"; \
# 		docker-compose up; \
# 	fi

# # Shutdown DB container
# docker-down:
# 	@if docker compose down 2>/dev/null; then \
# 		: ; \
# 	else \
# 		echo "Falling back to Docker Compose V1"; \
# 		docker-compose down; \
# 	fi


# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
