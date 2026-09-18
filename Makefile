.PHONY: build start stop restart logs logs-broker logs-frontend

# Build all service images
build:
	sudo docker compose build

# Build and start all services in the background
start:
	sudo docker compose up -d
	@echo "broker-service is running on port 8080"
	@echo "frontend-service is running on port 5173"

# Stop all running containers
stop:
	sudo docker compose down

# Restart all services
restart:
	sudo docker compose down
	sudo docker compose up -d

# View logs from all services
logs:
	sudo docker compose logs -f

# View logs from just the broker
logs-broker:
	sudo docker compose logs -f broker-service

# View logs from just the frontend
logs-frontend:
	sudo docker compose logs -f frontend-service