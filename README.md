# daily-shit
DailyShit is an example microservice newspaper.

* frontend:
    * website
* gateway
* microservices
    * user service
    * post service
    * comment service
    * newsletter service
    * ...
    
website <-> Gateway (graphql)  
gateway <-> microservices (gRPC)  
microservice <-> microservice (nats) 

database: mongo db

# Development requirements

install protoc:  
https://grpc.io/docs/languages/go/quickstart/

# Run

`go generate ./user-service .`  
`go run ./user-service .`  
`go run main.go`  