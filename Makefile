# Makefile for Fibonacci Web Service

# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOFMT=gofmt
BINARY_NAME=fibonacci-web-service
SRC_DIR=web-service
DOCKER_IMAGE_NAME=fibonacci-web-service

# Install dependencies
.PHONY: deps
deps:
	cd $(SRC_DIR) && $(GOGET) -v ./...

# Build the project
.PHONY: build
build:
	cd $(SRC_DIR) && $(GOBUILD) -o ../$(BINARY_NAME) -v

# Run tests
.PHONY: test
test:
	cd $(SRC_DIR) && $(GOTEST) -v ./...

# Clean build files
.PHONY: clean
clean:
	cd $(SRC_DIR) && $(GOCLEAN)
	rm -f $(BINARY_NAME)

# Install the project
.PHONY: install
install:
	cd $(SRC_DIR) && $(GOINSTALL) -v ./...

# Format the code
.PHONY: fmt
fmt:
	$(GOFMT) -w $(SRC_DIR)

# Build Docker image
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# Run Docker container
.PHONY: run-docker
run-docker: docker-build
	docker run -p 8080:8080 $(DOCKER_IMAGE_NAME)