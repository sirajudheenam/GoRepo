# GO from scratch

## Check the go version

```bash
> go version
go version go1.20.6 darwin/arm64
```

## Check the go path

```bash
> which go
/opt/homebrew/bin/go
```

## To build go

> go build .
> go: go.mod file not found in current directory or any parent directory; see 'go help modules'

### We have an error message

`go: go.mod file not found in current directory or any parent directory; see 'go help modules'`

### Solution:

```bash
go env -w GO111MODULE=off
```

## Look at go environment variables

```bash
go env

GO111MODULE="off"
GOARCH="arm64"
GOBIN=""
GOCACHE="${HOME}/Library/Caches/go-build"
GOENV="${HOME}/Library/Application Support/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="arm64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="${HOME}/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="${HOME}/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/opt/homebrew/Cellar/go/1.20.6/libexec"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/opt/homebrew/Cellar/go/1.20.6/libexec/pkg/tool/darwin_arm64"
GOVCS=""
GOVERSION="go1.20.6"
GCCGO="gccgo"
AR="ar"
CC="cc"
CXX="c++"
CGO_ENABLED="1"
GOMOD=""
GOWORK=""
CGO_CFLAGS="-O2 -g"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-O2 -g"
CGO_FFLAGS="-O2 -g"
CGO_LDFLAGS="-O2 -g"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch arm64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/3b/vx9r07dn7w74ys5zxn9lzpzc0000gp/T/go-build3679459993=/tmp/go-build -gno-record-gcc-switches -fno-common"

```

## You can also set your project as a part of a go module

```
go env -w GO111MODULE=auto
go mod init
go mod tidy
```

## build your code

```bash
cd gophercises/quiz
go build . && ./quiz -csv=abc.csv -limit=2
```

## Timer and Ticker functions

- Timer: sends a message to a channel once after 5 seconds.
- Ticker: sends a message to a channel every 5 seconds.

## Don't set $GOROOT, it will be set automatically

## Replacing modules

```bash
go mod edit -replace example.com/greetings=../greetings
go mod edit -replace=github.com/${YOUR_GIT_USERNAME}/GoRepo/cli/basics/cmd=./cmd
go mod edit -replace=github.com/${YOUR_GIT_USERNAME}/GoRepo/cli/basics/cmd/root=./cmd/root
go mod edit -replace=github.com/${YOUR_GIT_USERNAME}/GoRepo/cli/sarpamcli/cmd=./cmd
go mod edit -replace=github.com/${YOUR_GIT_USERNAME}/GoRepo/cli/sarpamcli/cmd/root=./cmd/root
```

## Set an go env

```bash
go env -w GO111MODULE=off
go env | grep GO111MODULE
```

# Clean Cache:

```bash
go clean -cache
```

## Golang Linting

```bash

golangci-lint run ./...
golangci-lint cache status
golangci-lint cache clean

# linting topic

go version

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

export PATH=$PATH:$(go env GOPATH)/bin

golangci-lint version

golangci-lint run --config .golangci-exceptions.yaml ./...


brew install act

act -P ubuntu-latest=catthehacker/ubuntu:full-latest

# add this line .actrc

act pull_request
#  or
act push

```

## Go Tests:

```bash

go test -v ./pkg/proxy
go test -v ./pkg/proxy -json | go tool test2json


go test -v -run ^TestThatExpirationIsEffective$ ./...

go test -v -run ^TestThatExpirationIsEffective$ gitrepo/cronus/cronus/pkg/taskqueue

go test -v -run ^TestThatExpirationIsEffective$


go clean -cache
go clean -modcache
go clean -testcache
go clean -cache -modcache -testcache -i

# to find out the dependencies of a package
go mod why -m golang.org/x/net
go mod why -m golang.org/x/crypto
go mod graph | grep golang.org/x/net
go list -deps ./... | grep x/net
go list -m all
go list -deps
go mod graph || grep pkg.go.dev/crypto/x509

github.wdf.sap.corp/cronus/billing-event-service github.com/Azure/go-amqp@v1.5.0
github.wdf.sap.corp/cronus/billing-event-service github.com/beorn7/perks@v1.0.1
github.wdf.sap.corp/cronus/billing-event-service github.com/cespare/xxhash/v2@v2.3.0
github.wdf.sap.corp/cronus/billing-event-service github.com/coder/websocket@v1.8.14
github.wdf.sap.corp/cronus/billing-event-service github.com/google/uuid@v1.6.0
github.wdf.sap.corp/cronus/billing-event-service github.com/jinzhu/inflection@v1.0.0
github.wdf.sap.corp/cronus/billing-event-service github.com/munnerz/goautoneg@v0.0.0-20191010083416-a7dc8b61c822
github.wdf.sap.corp/cronus/billing-event-service github.com/prometheus/client_golang@v1.23.2
github.wdf.sap.corp/cronus/billing-event-service github.com/prometheus/client_model@v0.6.2
github.wdf.sap.corp/cronus/billing-event-service github.com/prometheus/common@v0.66.1
github.wdf.sap.corp/cronus/billing-event-service github.com/prometheus/procfs@v0.16.1
github.wdf.sap.corp/cronus/billing-event-service github.com/puzpuzpuz/xsync/v3@v3.5.1
github.wdf.sap.corp/cronus/billing-event-service github.com/tmthrgd/go-hex@v0.0.0-20190904060850-447a3041c3bc
github.wdf.sap.corp/cronus/billing-event-service github.com/uptrace/bun@v1.2.16
github.wdf.sap.corp/cronus/billing-event-service github.com/uptrace/bun/dialect/pgdialect@v1.2.16
github.wdf.sap.corp/cronus/billing-event-service github.com/uptrace/bun/driver/pgdriver@v1.2.16
github.wdf.sap.corp/cronus/billing-event-service github.com/vmihailenco/msgpack/v5@v5.4.1
github.wdf.sap.corp/cronus/billing-event-service github.com/vmihailenco/tagparser/v2@v2.0.0
github.wdf.sap.corp/cronus/billing-event-service github.wdf.sap.corp/cronus/billing-event-client@v0.0.13
github.wdf.sap.corp/cronus/billing-event-service go@1.25
github.wdf.sap.corp/cronus/billing-event-service go.opentelemetry.io/otel@v1.38.0
github.wdf.sap.corp/cronus/billing-event-service go.opentelemetry.io/otel/trace@v1.38.0
github.wdf.sap.corp/cronus/billing-event-service go.yaml.in/yaml/v2@v2.4.2
github.wdf.sap.corp/cronus/billing-event-service golang.org/x/crypto@v0.47.0
github.wdf.sap.corp/cronus/billing-event-service golang.org/x/oauth2@v0.34.0
github.wdf.sap.corp/cronus/billing-event-service golang.org/x/sync@v0.19.0
github.wdf.sap.corp/cronus/billing-event-service golang.org/x/sys@v0.40.0
github.wdf.sap.corp/cronus/billing-event-service google.golang.org/protobuf@v1.36.10
github.wdf.sap.corp/cronus/billing-event-service mellium.im/sasl@v0.3.2
github.com/Azure/go-amqp@v1.5.0 github.com/fortytw2/leaktest@v1.3.0
github.com/Azure/go-amqp@v1.5.0 github.com/google/go-cmp@v0.5.9
github.com/Azure/go-amqp@v1.5.0 github.com/stretchr/testify@v1.8.4
github.com/Azure/go-amqp@v1.5.0 github.com/davecgh/go-spew@v1.1.1
github.com/Azure/go-amqp@v1.5.0 github.com/pmezard/go-difflib@v1.0.0
github.com/Azure/go-amqp@v1.5.0 gopkg.in/yaml.v3@v3.0.1
github.com/coder/websocket@v1.8.14 go@1.23
github.com/prometheus/client_golang@v1.23.2 github.com/beorn7/perks@v1.0.1
github.com/prometheus/client_golang@v1.23.2 github.com/cespare/xxhash/v2@v2.3.0
github.com/prometheus/client_golang@v1.23.2 github.com/google/go-cmp@v0.7.0
github.com/prometheus/client_golang@v1.23.2 github.com/json-iterator/go@v1.1.12
github.com/prometheus/client_golang@v1.23.2 github.com/klauspost/compress@v1.18.0
github.com/prometheus/client_golang@v1.23.2 github.com/kylelemons/godebug@v1.1.0
github.com/prometheus/client_golang@v1.23.2 github.com/prometheus/client_model@v0.6.2
github.com/prometheus/client_golang@v1.23.2 github.com/prometheus/common@v0.66.1
github.com/prometheus/client_golang@v1.23.2 github.com/prometheus/procfs@v0.16.1
github.com/prometheus/client_golang@v1.23.2 go.uber.org/goleak@v1.3.0
github.com/prometheus/client_golang@v1.23.2 golang.org/x/sys@v0.35.0
github.com/prometheus/client_golang@v1.23.2 google.golang.org/protobuf@v1.36.8
github.com/prometheus/client_golang@v1.23.2 github.com/jpillora/backoff@v1.0.0
github.com/prometheus/client_golang@v1.23.2 github.com/kr/pretty@v0.3.1
github.com/prometheus/client_golang@v1.23.2 github.com/modern-go/concurrent@v0.0.0-20180306012644-bacd9c7ef1dd
github.com/prometheus/client_golang@v1.23.2 github.com/modern-go/reflect2@v1.0.2
github.com/prometheus/client_golang@v1.23.2 github.com/munnerz/goautoneg@v0.0.0-20191010083416-a7dc8b61c822
github.com/prometheus/client_golang@v1.23.2 github.com/mwitkow/go-conntrack@v0.0.0-20190716064945-2f068394615f
github.com/prometheus/client_golang@v1.23.2 go.yaml.in/yaml/v2@v2.4.2
github.com/prometheus/client_golang@v1.23.2 golang.org/x/net@v0.43.0
github.com/prometheus/client_golang@v1.23.2 golang.org/x/oauth2@v0.30.0
github.com/prometheus/client_golang@v1.23.2 golang.org/x/text@v0.28.0
github.com/prometheus/client_golang@v1.23.2 go@1.23.0
github.com/prometheus/client_model@v0.6.2 google.golang.org/protobuf@v1.36.6
github.com/prometheus/client_model@v0.6.2 go@1.22.0
github.com/prometheus/common@v0.66.1 github.com/alecthomas/kingpin/v2@v2.4.0
github.com/prometheus/common@v0.66.1 github.com/google/go-cmp@v0.7.0
github.com/prometheus/common@v0.66.1 github.com/julienschmidt/httprouter@v1.3.0
github.com/prometheus/common@v0.66.1 github.com/munnerz/goautoneg@v0.0.0-20191010083416-a7dc8b61c822
github.com/prometheus/common@v0.66.1 github.com/mwitkow/go-conntrack@v0.0.0-20190716064945-2f068394615f
github.com/prometheus/common@v0.66.1 github.com/prometheus/client_model@v0.6.2
github.com/prometheus/common@v0.66.1 github.com/stretchr/testify@v1.11.1
github.com/prometheus/common@v0.66.1 go.yaml.in/yaml/v2@v2.4.2
github.com/prometheus/common@v0.66.1 golang.org/x/net@v0.43.0
github.com/prometheus/common@v0.66.1 golang.org/x/oauth2@v0.30.0
github.com/prometheus/common@v0.66.1 google.golang.org/protobuf@v1.36.8
github.com/prometheus/common@v0.66.1 github.com/alecthomas/units@v0.0.0-20211218093645-b94a6e3cc137
github.com/prometheus/common@v0.66.1 github.com/beorn7/perks@v1.0.1
github.com/prometheus/common@v0.66.1 github.com/cespare/xxhash/v2@v2.3.0
github.com/prometheus/common@v0.66.1 github.com/davecgh/go-spew@v1.1.1
github.com/prometheus/common@v0.66.1 github.com/jpillora/backoff@v1.0.0
github.com/prometheus/common@v0.66.1 github.com/pmezard/go-difflib@v1.0.0
github.com/prometheus/common@v0.66.1 github.com/prometheus/client_golang@v1.20.4
github.com/prometheus/common@v0.66.1 github.com/prometheus/procfs@v0.15.1
github.com/prometheus/common@v0.66.1 github.com/rogpeppe/go-internal@v1.10.0
github.com/prometheus/common@v0.66.1 github.com/xhit/go-str2duration/v2@v2.1.0
github.com/prometheus/common@v0.66.1 golang.org/x/sys@v0.35.0
github.com/prometheus/common@v0.66.1 golang.org/x/text@v0.28.0
github.com/prometheus/common@v0.66.1 gopkg.in/check.v1@v1.0.0-20201130134442-10cb98267c6c
github.com/prometheus/common@v0.66.1 gopkg.in/yaml.v3@v3.0.1
github.com/prometheus/common@v0.66.1 go@1.23.0
github.com/prometheus/procfs@v0.16.1 github.com/google/go-cmp@v0.7.0
github.com/prometheus/procfs@v0.16.1 golang.org/x/sync@v0.13.0
github.com/prometheus/procfs@v0.16.1 golang.org/x/sys@v0.32.0
github.com/prometheus/procfs@v0.16.1 go@1.23.0
github.com/uptrace/bun@v1.2.16 github.com/jinzhu/inflection@v1.0.0
github.com/uptrace/bun@v1.2.16 github.com/puzpuzpuz/xsync/v3@v3.5.1
github.com/uptrace/bun@v1.2.16 github.com/rs/zerolog@v1.34.0
github.com/uptrace/bun@v1.2.16 github.com/stretchr/testify@v1.8.1
github.com/uptrace/bun@v1.2.16 github.com/tmthrgd/go-hex@v0.0.0-20190904060850-447a3041c3bc
github.com/uptrace/bun@v1.2.16 github.com/vmihailenco/msgpack/v5@v5.4.1
github.com/uptrace/bun@v1.2.16 github.com/davecgh/go-spew@v1.1.1
github.com/uptrace/bun@v1.2.16 github.com/mattn/go-colorable@v0.1.14
github.com/uptrace/bun@v1.2.16 github.com/mattn/go-isatty@v0.0.20
github.com/uptrace/bun@v1.2.16 github.com/niemeyer/pretty@v0.0.0-20200227124842-a10e7caefd8e
github.com/uptrace/bun@v1.2.16 github.com/pmezard/go-difflib@v1.0.0
github.com/uptrace/bun@v1.2.16 github.com/vmihailenco/tagparser/v2@v2.0.0
github.com/uptrace/bun@v1.2.16 golang.org/x/sys@v0.38.0
github.com/uptrace/bun@v1.2.16 gopkg.in/check.v1@v1.0.0-20200227125254-8fa46927fb4f
github.com/uptrace/bun@v1.2.16 gopkg.in/yaml.v3@v3.0.1
github.com/uptrace/bun@v1.2.16 go@1.24.0
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/stretchr/testify@v1.8.1
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/uptrace/bun@v1.2.16
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/davecgh/go-spew@v1.1.1
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/jinzhu/inflection@v1.0.0
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/kr/text@v0.2.0
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/pmezard/go-difflib@v1.0.0
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/puzpuzpuz/xsync/v3@v3.5.1
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/tmthrgd/go-hex@v0.0.0-20190904060850-447a3041c3bc
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/vmihailenco/msgpack/v5@v5.4.1
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 github.com/vmihailenco/tagparser/v2@v2.0.0
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 golang.org/x/sys@v0.38.0
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 gopkg.in/yaml.v3@v3.0.1
github.com/uptrace/bun/dialect/pgdialect@v1.2.16 go@1.24.0
github.com/uptrace/bun/driver/pgdriver@v1.2.16 github.com/stretchr/testify@v1.11.1
github.com/uptrace/bun/driver/pgdriver@v1.2.16 github.com/uptrace/bun@v1.2.16
github.com/uptrace/bun/driver/pgdriver@v1.2.16 go.opentelemetry.io/otel@v1.38.0
github.com/uptrace/bun/driver/pgdriver@v1.2.16 go.opentelemetry.io/otel/trace@v1.38.0
github.com/uptrace/bun/driver/pgdriver@v1.2.16 mellium.im/sasl@v0.3.2
github.com/uptrace/bun/driver/pgdriver@v1.2.16 github.com/davecgh/go-spew@v1.1.1
github.com/uptrace/bun/driver/pgdriver@v1.2.16 github.com/jinzhu/inflection@v1.0.0
github.com/uptrace/bun/driver/pgdriver@v1.2.16 github.com/pmezard/go-difflib@v1.0.0
github.com/uptrace/bun/driver/pgdriver@v1.2.16 github.com/puzpuzpuz/xsync/v3@v3.5.1
github.com/uptrace/bun/driver/pgdriver@v1.2.16 github.com/tmthrgd/go-hex@v0.0.0-20190904060850-447a3041c3bc
github.com/uptrace/bun/driver/pgdriver@v1.2.16 github.com/vmihailenco/msgpack/v5@v5.4.1
github.com/uptrace/bun/driver/pgdriver@v1.2.16 github.com/vmihailenco/tagparser/v2@v2.0.0
github.com/uptrace/bun/driver/pgdriver@v1.2.16 golang.org/x/crypto@v0.45.0
github.com/uptrace/bun/driver/pgdriver@v1.2.16 golang.org/x/sys@v0.38.0
github.com/uptrace/bun/driver/pgdriver@v1.2.16 gopkg.in/yaml.v3@v3.0.1
github.com/uptrace/bun/driver/pgdriver@v1.2.16 go@1.24.0
github.com/vmihailenco/msgpack/v5@v5.4.1 github.com/stretchr/testify@v1.6.1
github.com/vmihailenco/msgpack/v5@v5.4.1 github.com/vmihailenco/tagparser/v2@v2.0.0
github.com/vmihailenco/msgpack/v5@v5.4.1 github.com/davecgh/go-spew@v1.1.0
github.com/vmihailenco/msgpack/v5@v5.4.1 github.com/pmezard/go-difflib@v1.0.0
github.com/vmihailenco/msgpack/v5@v5.4.1 gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/Azure/go-amqp@v1.4.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/coder/websocket@v1.8.13
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/google/uuid@v1.6.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/prometheus/client_golang@v1.22.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/stretchr/testify@v1.10.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 golang.org/x/oauth2@v0.29.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/beorn7/perks@v1.0.1
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/cespare/xxhash/v2@v2.3.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/davecgh/go-spew@v1.1.1
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/kr/text@v0.2.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/munnerz/goautoneg@v0.0.0-20191010083416-a7dc8b61c822
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/pmezard/go-difflib@v1.0.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/prometheus/client_model@v0.6.1
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/prometheus/common@v0.62.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 github.com/prometheus/procfs@v0.15.1
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 golang.org/x/sys@v0.32.0
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 google.golang.org/protobuf@v1.36.5
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 gopkg.in/yaml.v3@v3.0.1
github.wdf.sap.corp/cronus/billing-event-client@v0.0.13 go@1.23.0
go@1.25 toolchain@go1.25
go.opentelemetry.io/otel@v1.38.0 github.com/go-logr/logr@v1.4.3
go.opentelemetry.io/otel@v1.38.0 github.com/go-logr/stdr@v1.2.2
go.opentelemetry.io/otel@v1.38.0 github.com/google/go-cmp@v0.7.0
go.opentelemetry.io/otel@v1.38.0 github.com/stretchr/testify@v1.11.1
go.opentelemetry.io/otel@v1.38.0 go.opentelemetry.io/auto/sdk@v1.1.0
go.opentelemetry.io/otel@v1.38.0 go.opentelemetry.io/otel/metric@v1.38.0
go.opentelemetry.io/otel@v1.38.0 go.opentelemetry.io/otel/trace@v1.38.0
go.opentelemetry.io/otel@v1.38.0 github.com/davecgh/go-spew@v1.1.1
go.opentelemetry.io/otel@v1.38.0 github.com/kr/text@v0.2.0
go.opentelemetry.io/otel@v1.38.0 github.com/pmezard/go-difflib@v1.0.0
go.opentelemetry.io/otel@v1.38.0 gopkg.in/yaml.v3@v3.0.1
go.opentelemetry.io/otel@v1.38.0 go@1.23.0
go.opentelemetry.io/otel/trace@v1.38.0 github.com/google/go-cmp@v0.7.0
go.opentelemetry.io/otel/trace@v1.38.0 github.com/stretchr/testify@v1.11.1
go.opentelemetry.io/otel/trace@v1.38.0 go.opentelemetry.io/otel@v1.38.0
go.opentelemetry.io/otel/trace@v1.38.0 github.com/davecgh/go-spew@v1.1.1
go.opentelemetry.io/otel/trace@v1.38.0 github.com/pmezard/go-difflib@v1.0.0
go.opentelemetry.io/otel/trace@v1.38.0 gopkg.in/yaml.v3@v3.0.1
go.opentelemetry.io/otel/trace@v1.38.0 go@1.23.0
go.yaml.in/yaml/v2@v2.4.2 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
golang.org/x/crypto@v0.47.0 golang.org/x/net@v0.48.0
golang.org/x/crypto@v0.47.0 golang.org/x/sys@v0.40.0
golang.org/x/crypto@v0.47.0 golang.org/x/term@v0.39.0
golang.org/x/crypto@v0.47.0 golang.org/x/text@v0.33.0
golang.org/x/crypto@v0.47.0 go@1.24.0
golang.org/x/oauth2@v0.34.0 cloud.google.com/go/compute/metadata@v0.3.0
golang.org/x/oauth2@v0.34.0 go@1.24.0
golang.org/x/sync@v0.19.0 go@1.24.0
golang.org/x/sys@v0.40.0 go@1.24.0
google.golang.org/protobuf@v1.36.10 github.com/golang/protobuf@v1.5.0
google.golang.org/protobuf@v1.36.10 github.com/google/go-cmp@v0.7.0
google.golang.org/protobuf@v1.36.10 go@1.23
mellium.im/sasl@v0.3.2 golang.org/x/crypto@v0.27.0


# Update all
go get -u ./...

# Routine:
go mod tidy
go mod vendor
go test ./...
golangci-lint run
govulncheck ./...


```

References:

[Nice Cheat Sheet](https://github.com/a8m/golang-cheat-sheet)
