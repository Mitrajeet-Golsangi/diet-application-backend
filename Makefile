# Variables used for this Makefile
MAIN_FILE=cmd/main.go
OUTPUT_DIR=build/output

default: help

# Running the application locally
.PHONY: run
run:
	go run $(MAIN_FILE)

# Testing the application
.PHONY: test
test:
	go test

# Organize the dependencies for the application and update the vendor data
.PHONY: dep
dep:
	go mod tidy && go mod vendor

# Build the application for the production
.PHONY: build
build:
	go build -o $(OUTPUT_DIR)/main $(MAIN_FILE)

# Clean the unnecessary files
.PHONY: clean
clean:
	rm -rf *.exe *.exe~ *.dll *.so *.dylib *.out $(OUTPUT_DIR)

# Showing the help message to run different commands for the application
.PHONY: help
help:
	@echo "make run\t\t- Run the application server in development mode"
	@echo "make help\t\t- Show this help message"
	@echo "make test\t\t- Run the tests for the API"
	@echo "make dep\t\t- Organize the dependencies for the application and update the vendor data"
	@echo "make build\t\t- Build the application"
	@echo "make clean\t\t- Clean the application's build files and other generated files"
	@echo "make docker-build\t- Build the docker image for the application"
	@echo "make docker-run\t\t- Run the docker container for the application"
