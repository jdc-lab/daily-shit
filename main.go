// file main.go is just for testing the microservices without full implementation while developing.
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "./user-service/proto/user"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":4040", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()
	c := pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	// create user
	u := pb.CreateUserRequest{
		Username: "aligator",
		Email:    "aligator@suncraft-server.de",
		Password: "superpassword",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	createReq, err := c.Create(ctx, &u)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("user id: %s", createReq.GetId())

	authReq, err := c.Auth(ctx, &pb.AuthRequest{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		log.Fatalf("could not authenticate user: %v", err)
	}
	log.Printf("auth: %v", authReq)

	claims, err := c.ValidateToken(ctx, &pb.ValidateTokenRequest{
		Token: authReq.Token,
	})
	if err != nil {
		log.Fatalf("could not validate token: %v", err)
	}
	log.Printf("validate: %v", claims)

	getReq, err := c.Get(ctx, &pb.GetUserRequest{Id: createReq.Id})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("user: %v", getReq)

	// this should only work for the first run, as in the second run the first created user is no admin
	u.IsAdmin = true
	u.Claims = claims
	createReq, err = c.Create(ctx, &u)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("user id: %s", createReq.GetId())
}
