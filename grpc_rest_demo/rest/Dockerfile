# REST Dockerfile
FROM golang:1.25.3-alpine3.22 AS builder
WORKDIR /app
COPY rest_server.go .
RUN go mod init rest && go build -o rest_server .

FROM gcr.io/distroless/base
COPY --from=builder /app/rest_server /rest_server
EXPOSE 8080
ENTRYPOINT ["/rest_server"]
