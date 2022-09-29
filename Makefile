PKG_LIST := $(shell go list ./... | grep -v /vendor/)
CURRENT_DIR=$(shell pwd)
APP_CMD_DIR=${CURRENT_DIR}/cmd
APP=$(shell basename ${CURRENT_DIR})


run:
	go run cmd/main.go

swag-gen:
	echo ${REGISTRY}
	swag init -g cmd/main.go -o api/docs 

lint:
	golint -set_exit_status ${PKG_LIST}

unit-tests: ## Run unit-tests
	go test -mod=vendor -v -cover -short ${PKG_LIST}

race: ## Run data race detector
	go test -mod=vendor -race -short ${PKG_LIST}

msan: ## Run memory sanitizer. If this test fails, you need to write the following command: export CC=clang (if you have installed clang)
	env CC=clang env CXX=clang++ go test -mod=vendor -msan -short ${PKG_LIST}

build:
	CGO_ENABLED=1 GOOS=linux go build -mod=vendor -a -installsuffix cgo -race -ldflags "-extldflags '-static'" -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

set-env:
	./scripts/set-env.sh ${CURRENT_DIR}

vendor: 
	go mod vendor

tidy:
	go mod tidy

create-env: ## Creates .env file with example .env file
	cp ./.env.example ./.env

.PHONY: vendor