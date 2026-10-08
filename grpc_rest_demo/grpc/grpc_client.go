package main

import (
	"context"
	"log"
	"time"

	// pb "example.com/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: "42"})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %s - %s", resp.Id, resp.Name)
}
