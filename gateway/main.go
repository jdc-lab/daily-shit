//go:generate go run github.com/99designs/gqlgen generate
package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jdc-lab/daily-shit/common"
	"github.com/jdc-lab/daily-shit/gateway/graph"
	"github.com/jdc-lab/daily-shit/gateway/graph/generated"
	pb "github.com/jdc-lab/daily-shit/user-service/proto/user"
	"google.golang.org/grpc"
)

const defaultPort = "8080"

func main() {
	service := common.GetRandomServiceWithConsul("user-service")
	if service == nil {
		panic("no user-service found")
	}

	host := net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))
	// Set up a connection to the server.
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()
	userService := pb.NewUserServiceClient(conn)

	// ToDo: register also with consul
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			UserService: userService,
		},
		Directives: graph.Directives(),
	}))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(Authenticator(userService))

	r.Handle("/", playground.Handler("GraphQL playground", "/v1"))
	r.Handle("/v1", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
