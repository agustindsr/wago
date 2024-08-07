# Variables
APP_NAME := myproject
GOOS := linux
GOARCH := amd64
ZIP_FILE := $(APP_NAME).zip
LAMBDA_FUNCTION_NAME := my-lambda-function
WASM_OUTPUT := play.wasm

# Default target
all: build_lambda optimize zip deploy

# Build the Go application for AWS Lambda
build_lambda:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o main main.go

# Package the build output and static files into a zip file
zip: build_lambda optimize
	zip -r $(ZIP_FILE) main docs/ $(WASM_OUTPUT)

# Deploy the zip file to AWS Lambda
deploy: zip
	aws lambda update-function-code --function-name $(LAMBDA_FUNCTION_NAME) --zip-file fileb://$(ZIP_FILE)

# Build the Go application for WebAssembly
build_wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

# Optimize the WebAssembly file
optimize: build_wasm
	wasm-opt main.wasm --enable-bulk-memory -Oz -o $(WASM_OUTPUT)

# Clean up build artifacts
clean:
	rm -f main main.wasm $(WASM_OUTPUT) $(ZIP_FILE)

# Help target
help:
	@echo "Makefile commands:"
	@echo "  make build_lambda      - Build the Go application for AWS Lambda"
	@echo "  make zip               - Package the build output and static files into a zip file"
	@echo "  make deploy            - Deploy the zip file to AWS Lambda"
	@echo "  make build_wasm        - Build the Go application for WebAssembly"
	@echo "  make optimize          - Optimize the WebAssembly file"
	@echo "  make clean             - Clean up build artifacts"
	@echo "  make all               - Build, package, and deploy the application"

# Regla por defecto que ejecuta ambas tareas
build: build_lambda optimize

# Clean up build artifacts
clean:
	rm -f main main.wasm $(WASM_OUTPUT) $(ZIP_FILE)
