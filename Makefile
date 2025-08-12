SHELL := /bin/bash

APP     := server
CMD     := ./cmd/server
TRANS   ?= ./transactions.json
MODULE  ?= assesment

.PHONY: help init run test

help:
	@echo "Targets:"
	@echo "  make init        - инициализировать go.mod (module=$(MODULE)) и tidy"
	@echo "  make run         - run server with --transactions=$(TRANS)"
	@echo "  make test        - запустить тесты"

init:
	@if [ ! -f go.mod ]; then \
	  echo ">> go mod init $(MODULE)"; \
	  go mod init $(MODULE); \
	fi
	@echo ">> go mod tidy"
	@go mod tidy

run:
	@echo ">> go run $(CMD) --transactions $(TRANS)"
	@go run $(CMD) --transactions $(TRANS)


test:
	@go test ./...