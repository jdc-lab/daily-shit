// file main.go is just for testing the microservices without full implementation while developing.
package main

import (
	"log"

	"google.golang.org/grpc"

	pb "github.com/jdc-lab/daily-shit/user-service/proto/user"
)

func main() {
	// Todo

	// Set up a connection to the server.
	conn, err := grpc.Dial(":8100", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()
	_ = pb.NewUserServiceClient(conn)
}
