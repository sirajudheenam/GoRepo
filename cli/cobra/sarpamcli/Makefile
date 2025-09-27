PKG:=github.com/sirajudheenam/GoRepo/cli/cobra/sarpamcli
APP_NAME:=sarpam
PWD:=$(shell pwd)
UID:=$(shell id -u)
VERSION:=$(shell git describe --tags --always --dirty="-dev")
LDFLAGS:=-X main.Version=$(VERSION) -w -s

export GOPRIVATE:=github.com/sirajudheenam
export CGO_ENABLED:=0

.PHONY: docker

build: fmt vet linux darwin windows

linux:
	GOOS=linux GOARCH=amd64 go build -mod=vendor -trimpath -ldflags="$(LDFLAGS)" -o bin/sarpamcli .

darwin:
	GOOS=darwin GOARCH=amd64 go build -mod=vendor -trimpath -ldflags="$(LDFLAGS)" -o bin/sarpamcli_darwin .

windows:
	GOOS=windows GOARCH=amd64 go build -mod=vendor -trimpath -ldflags="$(LDFLAGS)" -o bin/sarpamcli.exe .

docker:
	docker pull golang:latest
	docker run -ti --rm -e GOCACHE=/tmp -v $(PWD):/$(APP_NAME) -u $(UID):$(UID) --workdir /$(APP_NAME) golang:latest make build

fmt:
	gofmt -s -w sarpamcli internal

vet:
	go vet -mod=vendor . ./internal/...

static:
	staticcheck . ./internal/...

mod:
	go mod vendor

test:
	go test -v -count 1 --short . ./internal/...

