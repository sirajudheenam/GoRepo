# Concurrency

Pre-requisites:

[Apache Bench](https://httpd.apache.org/docs/2.4/programs/ab.html)

```bash
go mod init github.com/GoRepo/concurrency/simple
go mod tidy
go run main.go

brew install httpd  #(or) brew install apr-util
which ab
# /opt/homebrew/bin/ab

# Now on another terminal 

# Test normal cURL call
curl "localhost:8000/set?name=x&val=0"
ok
curl "localhost:8000/get?name=x"
x: 0
curl "localhost:8000/inc?name=x"
ok
curl "localhost:8000/get?name=x"
x: 1


# do the benchmark with ab 

ab -n 20000 -c 200 "127.0.0.1:8000/inc?name=i"

# it looks like this on server output

# goroutine 1246 [IO wait]:
# internal/poll.runtime_pollWait(0x127400600, 0x72)
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/netpoll.go:351 +0xa0
# internal/poll.(*pollDesc).wait(0x14000174b00?, 0x140002112e1?, 0x0)
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/poll/fd_poll_runtime.go:84 +0x28
# internal/poll.(*pollDesc).waitRead(...)
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/poll/fd_poll_runtime.go:89
# internal/poll.(*FD).Read(0x14000174b00, {0x140002112e1, 0x1, 0x1})
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/poll/fd_unix.go:165 +0x1e0
# net.(*netFD).Read(0x14000174b00, {0x140002112e1?, 0x10057f0f0?, 0x14000316694?})
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/net/fd_posix.go:68 +0x28
# net.(*conn).Read(0x140002a4068, {0x140002112e1?, 0x0?, 0x0?})
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/net/net.go:196 +0x34
# net/http.(*connReader).backgroundRead(0x140002112c0)
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/net/http/server.go:702 +0x38
# created by net/http.(*connReader).startBackgroundRead in goroutine 1411
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/net/http/server.go:698 +0xb8

# goroutine 1302 [IO wait]:
# internal/poll.runtime_pollWait(0x127400c00, 0x72)
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/netpoll.go:351 +0xa0
# internal/poll.(*pollDesc).wait(0x14000174200?, 0x140002b8aa1?, 0x0)
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/poll/fd_poll_runtime.go:84 +0x28
# internal/poll.(*pollDesc).waitRead(...)
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/poll/fd_poll_runtime.go:89
# internal/poll.(*FD).Read(0x14000174200, {0x140002b8aa1, 0x1, 0x1})
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/poll/fd_unix.go:165 +0x1e0
# net.(*netFD).Read(0x14000174200, {0x140002b8aa1?, 0x10057f0f0?, 0x14000989594?})
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/net/fd_posix.go:68 +0x28
# net.(*conn).Read(0x140002a4050, {0x140002b8aa1?, 0x0?, 0x0?})
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/net/net.go:196 +0x34
# net/http.(*connReader).backgroundRead(0x140002b8a80)
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/net/http/server.go:702 +0x38
# created by net/http.(*connReader).startBackgroundRead in goroutine 1008
#         /opt/homebrew/Cellar/go/1.25.1/libexec/src/net/http/server.go:698 +0xb8

# now run with mutex enabled 

cd ../mutex_server/ 
go run main.go

# do the benchmark with ab on another terminal

ab -n 20000 -c 200 "127.0.0.1:8000/inc?name=i"

# Connection Times (ms)
#               min  mean[+/-sd] median   max
# Connect:        0    6   7.2      5      97
# Processing:     0   10   6.5      7      38
# Waiting:        0   10   6.4      6      36
# Total:          0   16  10.2     13     116

# Percentage of the requests served within a certain time (ms)
#   50%     13
#   66%     19
#   75%     20
#   80%     21
#   90%     24
#   95%     37
#   98%     50
#   99%     54
#  100%    116 (longest request)

# concurrency problem is solved, however 

```

