.PHONY: default run build fmt help docker-build

# Variables
APP_NAME := "<%= projectName %>"
MAIN := "cmd/main.go"

default: run

run:
	@echo "Running the application..."
	go run cmd/main.go

clean:
	@echo "Cleaning the application..."
	@rm -rf bin

build: clean
	@echo "Building the application..."
	go build -o bin/${APP_NAME} ${MAIN}

build-release: clean
	@echo "Building the application..."
	CGO_ENABLED=0 GOOS=linux go build -o bin/${APP_NAME} ${MAIN}

fmt:
	@echo "Formatting the application..."
	go fmt ./...

docker-build:
	@echo "Building the Docker image..."
	docker build -t ${APP_NAME} .

docker-run:
	@echo "Running the Docker container..."
	docker run -e PORT=8080 -p 3001:8080 --rm ${APP_NAME}

docker-clean:
	@echo "Cleaning the Docker container..."
	docker rmi -f ${APP_NAME}

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run    - Run the application"
	@echo "  build  - Build the application"
	@echo "  help   - Display this help message"
