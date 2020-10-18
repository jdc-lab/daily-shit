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

* As this repo is currently private, you need to get a github access token and use it for this repo:  
` git config --global url."https://$(YOUR_NAME):$(YOUR_TOKEN)@github.com/jdc-lab/daily-shit".insteadOf "https://github.com/jdc-lab/daily-shit"`  
You get the token from Github -> settings -> Developer settings -> Personal access tokens -> Generate new token -> check at least the "repo" scope  
__Or alternatively__: if you use already an ssh key:
`git config --global url."git@github.com/jdc-lab/daily-shit".insteadOf "https://github.com/jdc-lab/daily-shit"`
* Also you have to set the environment variable:  
`GOPRIVATE=github.com/jdc-lab/daily-shit` (maybe also in your IDE)  
and for me I had problems with GOSUMDB, that's why I set it to `GOSUMDB=off` for this project. (you may not need this)

* install protoc for go:  
https://grpc.io/docs/languages/go/quickstart/
* install gqlgen `go get github.com/99designs/gqlgen`
* install docker
* install docker-compose
* install make

# Run
## consul
To just run the consul container:  
`make consul-start` 
Services will be able to register to it using `localhost:8500` (default setting for the services) even if they do not run in docker.  

You can access the consul ui by `http://localhost:8500/ui/`

## local
This starts the services without docker. Note that they still need a consul instance which can still be started through docker (see above).  

__User service:__
* optionally if you changed the proto definitions: `go generate ./user-service .`   
* `make user-service`  