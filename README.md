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
