//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user/user.proto
package main

import (
	"flag"
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
		log.Fatal("empty secret")
	}

	listener, err := net.Listen("tcp", port())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, &handler{
		repo: &inMemoryRepository{},
		auth: &jwtAuthenticator{[]byte(*secret)},
	})

	registerServiceWithConsul("user-service")

	log.Println("setup finished - starting service")
	if e := srv.Serve(listener); e != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
