# Default variables (can be overridden via command line)
# Example: make YEAR=2024 DAY=day05 run
YEAR ?= 2025
DAY  ?= 1
FORMATTED_DAY := $(shell printf "%02d" $(DAY))
DIR := years/$(YEAR)/day$(FORMATTED_DAY)
ENTRYPOINT := $(DIR)/day$(FORMATTED_DAY).go

.PHONY: run test_verbose test

# Default target (runs when you just type 'make')
all: run

# Run the main program
run:
	@echo ">>> Running $(ENTRYPOINT)"
	@go run $(ENTRYPOINT)

# Run tests with verbose output
test_verbose:
	@echo ">>> Testing $(DIR)"
	@go test -v $(DIR)/*.go

# Run tests
test:
	@echo ">>> Testing $(DIR)"
	@go test $(DIR)/*.go
