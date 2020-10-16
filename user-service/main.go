//go:generate protoc -I=proto/user --go_out=plugins=grpc:proto/user proto/user/user.proto
package main

import (
	"flag"
	"log"
	"net"
	"os"

	pb "./proto/user"
	"google.golang.org/grpc"
)

func main() {
	secret := flag.String("jwt-secret", "", "the jwt secret to be used, can also be provided using the environment variable 'JWT_SECRET'")
	flag.Parse()

	if secret == nil || *secret == "" {
		envSecret := os.Getenv("JWT_SECRTET")
		secret = &envSecret
	}

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, &handler{
		repo: &inMemoryRepository{},
		auth: &jwtAuthenticator{[]byte(*secret)},
	})

	if e := srv.Serve(listener); e != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
