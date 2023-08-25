# Variables
APP_NAME = reservify
MAIN_FILE = ./cmd/application/main.go
API_FILE = ./internal/adapters/delivery/http/api.go
DELIVERY_DIR = ./internal/adapters/delivery/docs

# Commands
GO = go
GO_RUN = $(GO) run
GO_BUILD = $(GO) build
GO_GENERATE_DOCS = swag init -g $(API_FILE) -o $(DELIVERY_DIR)

# Build
.PHONY: run build docs clean

# Generate
run:
	$(GO_RUN) $(MAIN_FILE)

build:
	$(GO_BUILD) -o $(APP_NAME) $(MAIN_FILE)

docs:
	$(GO_GENERATE_DOCS)

clean:
	rm -f $(APP_NAME)

default: docs run
