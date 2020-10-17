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

# Permissions

The idea is to check the jwt token for each request to the gateway automatically
and then just pass the received claims from the gateway to each subsequent microservice.  
So the actual permissions are only validated one time -> less overhead  
The microservices are still able to decide themselves if they allow or deny something
by checking the claims.

Note that this means that there is no security at all if the gateway is bypassed
so from the outside everything should always go through the gateway.

# Development requirements

install protoc:  
https://grpc.io/docs/languages/go/quickstart/

# Run

`go generate ./user-service .`  
`go run ./user-service -jwt-secret="your-super-secret"`  
`go run main.go`  