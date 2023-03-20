.PHONY: build test check docker/build docker/up docker/down help

IMAGE_NAME := rfdez/diade

all: help

## Build the project
build:
	@CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo --ldflags="-s -w" -o diade ./cmd/diade-api/main.go

## Run the project tests
test:
	@go test -v -cover ./...

## Check the project code style
check:
	@golangci-lint run

## Build the docker image
docker/build:
	@docker build -t $(IMAGE_NAME) .

## Create and start containers
docker/up:
	@docker compose --compatibility up --build -d

## Stop and remove containers, networks and, optionally images and volumes
docker/down:
	@docker compose --compatibility down --remove-orphans -v

## This help screen
help:
	@printf "Available targets:\n\n"
	@awk '/^[a-zA-Z\-\_0-9%:\\]+/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
		helpCommand = $$1; \
		helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
	gsub("\\\\", "", helpCommand); \
	gsub(":+$$", "", helpCommand); \
		printf "  \x1b[32;01m%-35s\x1b[0m %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"
