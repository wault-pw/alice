-include .env.make

.DEFAULT_GOAL := help
BIN:=$(CURDIR)/bin
BIN_LINTER:=$(BIN)/golangci-lint
VERSION:=$(shell cat VERSION)
REGISTRY_DOMAIN=ghcr.io
REGISTRY_NAME=ghcr.io/oka-is/alice
PG_DSN:=postgres://localhost:5432/alice?sslmode=disable&timezone=utc

help:
	@echo 'Available targets: $(VERSION)'
	@echo ' '
	@echo '  make db:status'
	@echo '  make db:up'
	@echo '  make db:down'
	@echo '  make NAME="create_users" db:create'

.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(BIN_LINTER)),)
	$(info Downloading golangci-lint)
	GOBIN=$(BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0
endif

db\:up:
	goose -dir migrations postgres "$(PG_DSN)" up

db\:down:
	goose -dir migrations postgres "$(PG_DSN)" down

db\:create: NAME=$NAME
db\:create:
	goose -dir migrations postgres "$(PG_DSN)" create $(NAME) sql

proto:
	protoc --proto_path=protos --go_out=. alice_v1.proto

test:
	go test -count=1 -p 4 -race -cover -covermode atomic ./...

test\:wasm:
	yarn --cwd js run test

lint: install-lint
lint:
	$(BIN_LINTER) run --config=.golangci.yaml ./...

generate:
	go generate ./...

generate_build:
	go generate cmd/goose.go

build\:wasm:
	go generate wasm.go
	GOOS=js GOARCH=wasm CGO_ENABLED=0 go build -a -installsuffix cgo -o build/alice.wasm wasm.go

build\:linux: generate_build
build\:linux:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -X 'main.Version=$(VERSION)'" -a -installsuffix cgo -o build/linux

build\:mac: generate_build
build\:mac:
	env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -X 'main.Version=$(VERSION)'" -a -installsuffix cgo -o build/mac

docker\:build: export TAG=$(VERSION)
docker\:build:
	docker build --no-cache -f ./Dockerfile -t ${REGISTRY_NAME}:${TAG} .
	docker tag ${REGISTRY_NAME}:${TAG} ${REGISTRY_NAME}:latest

docker\:push: export TAG=$(VERSION)
docker\:push:
	@echo $(REGISTRY_PASSWORD) | docker login ${REGISTRY_DOMAIN} --username ${REGISTRY_USERNAME} --password-stdin
	docker push ${REGISTRY_NAME}:${TAG}
	docker push ${REGISTRY_NAME}:latest
