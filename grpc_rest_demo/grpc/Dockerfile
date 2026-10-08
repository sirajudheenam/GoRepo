# gRPC Dockerfile
FROM golang:1.25.3-alpine3.22 AS builder
WORKDIR /app
COPY user.proto grpc_server.go go.mod go.sum ./
RUN go mod tidy && \
    protoc --go_out=. --go-grpc_out=. user.proto && \
    go build -o grpc_server .

FROM gcr.io/distroless/base
COPY --from=builder /app/grpc_server /grpc_server
EXPOSE 50051
ENTRYPOINT ["/grpc_server"]