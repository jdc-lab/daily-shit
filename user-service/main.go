//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user/user.proto
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "daily-shit/user-service/proto/user"

	"google.golang.org/grpc"
)

func main() {
	secret := flag.String("jwt-secret", "", "the jwt secret to be used, can also be provided using the environment variable 'JWT_SECRET'")

	flag.Parse()

	if *secret == "" {
		envSecret := os.Getenv("JWT_SECRET")
		secret = &envSecret
	}

	if *secret == "" {
		panic("empty secret")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 4040))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	srv := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(srv, &handler{
		repo: &inMemoryRepository{},
		auth: &jwtAuthenticator{[]byte(*secret)},
	})

	if e := srv.Serve(listener); e != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
