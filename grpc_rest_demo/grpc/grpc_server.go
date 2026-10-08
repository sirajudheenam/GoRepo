package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "example.com/user"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		Id:   req.Id,
		Name: fmt.Sprintf("User %s", req.Id),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("gRPC server running on :50051")
	log.Fatal(grpcServer.Serve(lis))
}
