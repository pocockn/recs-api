PACKAGES ?= "./..."
REPONAME := "recs-api"

DEFAULT: test

build:
	@GO111MODULE=on go build "${PACKAGES}"

build-image:
	@docker build -t pocockn/${REPONAME} .

install:
	@echo "=> Install dependencies"
	@GO111MODULE=on go mod download

run:
	@go build -ldflags "-X main.Version=dev"
	@ENV=development AWS_REGION=eu-west-1 ./recs-api

test:
	@GO111MODULE=on go test "${PACKAGES}" -cover

vet:
	@@GO111MODULE=on go vet "${PACKAGES}"