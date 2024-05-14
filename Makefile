-include .env
export

CURRENT_DIR=$(shell pwd)
APP=evrone_api_gateway
CMD_DIR=./cmd

#SERVER_HOST = localhost
#SERVER_PORT = 9050
#
#REDIS_HOST = localhost
#REDIS_PORT = 6379
#
#BOOKING_SERVICE_GRPC_HOST = localhost
#BOOKING_SERVICE_GRPC_PORT = 9090
#
#HEALTHCARE_SERVICE_GRPC_HOST = localhost
#HEALTHCARE_SERVICE_GRPC_PORT = 9080
#
#SESSION_SERVICE_GRPC_HOST = localhost
#SESSION_SERVICE_GRPC_PORT = 9060
#
#USER_SERVICE_GRPC_HOST = localhost
#USER_SERVICE_GRPC_PORT = 9070
#
#OTLP_COLLECTOR_HOST = localhost
#
#MINIO_SERVICE_ENDPOINT = localhost:9000
#
#POSTGRES_USER = postgres
#POSTGRES_PASSWORD = 20030505
#POSTGRES_HOST = localhost
#POSTGRES_PORT = 5432
#POSTGRES_DATABASE = dennic

# -------------DOCKER--------------

SERVER_HOST = dennic_api_gateway
SERVER_PORT = 9050

REDIS_HOST = redisdb
REDIS_PORT = 6379

BOOKING_SERVICE_GRPC_HOST = dennic_booking_service
BOOKING_SERVICE_GRPC_PORT = 9090

HEALTHCARE_SERVICE_GRPC_HOST = dennic_healthcare_service
HEALTHCARE_SERVICE_GRPC_PORT = 9080

SESSION_SERVICE_GRPC_HOST = dennic_session_service
SESSION_SERVICE_GRPC_PORT = 9060

USER_SERVICE_GRPC_HOST = dennic_user_service
USER_SERVICE_GRPC_PORT = 9070

OTLP_COLLECTOR_HOST = otlp-collector

MINIO_SERVICE_ENDPOINT = minio:9000

POSTGRES_USER = postgres
POSTGRES_PASSWORD = 20030505
POSTGRES_HOST = postgresdb
POSTGRES_PORT = 5432
POSTGRES_DATABASE = dennic



# build for current os
.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

# build for linux amd64
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOARCH="amd64" GOOS=linux go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

# run service
.PHONY: run
run:
	go run ${CMD_DIR}/app/main.go

# proto
.PHONY: proto-gen
proto-gen:
	./scripts/gen-proto.sh


# go generate
.PHONY: go-gen
go-gen:
	go generate ./...

# generate swagger
.PHONY: swagger-gen
swagger-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs

# run test
.PHONY: test
test:
	go test -v -cover -race ./internal/...

# migrate
.PHONY: migrate
migrate:
	migrate -source file://migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable up

# -------------- for deploy --------------
build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

.PHONY: pull-proto-module
pull-proto-module:
	git submodule update --init --recursive

.PHONY: update-proto-module
update-proto-module:
	git submodule update --remote --merge