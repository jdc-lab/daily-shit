//go:generate protoc -I=proto/user --go_out=plugins=grpc:proto/user proto/user/user.proto
package main

import (
	"log"
	"net"

	pb "./proto/user"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, &handler{
		repo: &inMemoryRepository{},
	})
	//reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
