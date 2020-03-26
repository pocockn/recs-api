BINARY_UNIX=$(REPONAME)_unix
PACKAGES ?= "./..."
DOCKERNAME = "pococknick91"
REPONAME ?= "recs-api"
IMG ?= ${DOCKERNAME}/${REPONAME}:${VERSION}
LATEST ?= ${DOCKERNAME}/${REPONAME}:latest
VERSION = $(shell cat ./VERSION)

DEFAULT: test

build:
	@GO111MODULE=on go build "${PACKAGES}"

build-image:
	@docker build -t ${IMG} .

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_UNIX) -v

install:
	@echo "=> Install dependencies"
	@GO111MODULE=on go mod download

push-to-registry:
	@docker login -u ${DOCKER_USER} -p ${DOCKER_PASS}
	@docker build -t ${IMG} .
	@docker tag ${IMG} ${LATEST}
	echo "=> Pushing ${IMG} & ${LATEST} to docker"
	@docker push ${DOCKERNAME}/${REPONAME}

run:
	@go build -ldflags "-X main.Version=dev"
	@ENV=development AWS_REGION=eu-west-1 ./recs-api

test:
	@GO111MODULE=on go test "${PACKAGES}" -cover

vet:
	@@GO111MODULE=on go vet "${PACKAGES}"