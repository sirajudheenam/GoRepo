PKG:=github.com/sirajudheenam/GoRepo/k8s-client
APP_NAME:=k8s-client
PWD:=$(shell pwd)
UID:=$(shell id -u)
VERSION:=$(shell git describe --tags --always --dirty="-dev")
LDFLAGS:=-X main.Version=$(VERSION) -w -s
BUILDIMAGE:=golang:1.25-alpine

export CGO_ENABLED:=0

.PHONY: docker

build: fmt vet static test linux darwin windows

linux:
	GOOS=linux GOARCH=amd64 go build -mod=vendor -trimpath -ldflags="$(LDFLAGS)" -o bin/k8s-client ./in-cluster

darwin:
	GOOS=darwin GOARCH=amd64 go build -mod=vendor -trimpath -ldflags="$(LDFLAGS)" -o bin/k8s-client_darwin ./in-cluster

windows:
	GOOS=windows GOARCH=amd64 go build -mod=vendor -trimpath -ldflags="$(LDFLAGS)" -o bin/k8s-client.exe ./in-cluster

docker:
	docker pull golang:1.25-alpine
	docker run -ti --rm -e GOCACHE=/tmp -v $(PWD):/$(APP_NAME) -u $(UID):$(UID) --workdir /$(APP_NAME) $(BUILDIMAGE) make build

fmt:
	gofmt -s -w in-cluster

vet:
	go vet -mod=vendor ./in-cluster

static:
	staticcheck ./in-cluster

mod:
	go mod vendor

test:
	go test -v -count 1 --short ./in-cluster

