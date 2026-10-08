# gRPC Demo

```bash
# Step 1. Create user.proto

# Step 2. Generate Go code

# You’ll need protoc and the Go plugins installed:
protoc --go_out=. --go-grpc_out=. user.proto

# Run server:
go run grpc_server.go

# Run client:
go run grpc_client.go


docker build -t grpc-server .
docker run -p 50051:50051 grpc-server

kubectl apply -f grpc-deployment.yaml


```