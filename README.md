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

# gateway

The gateway is currently using graphql (just because I wanted to experiment with it).  
It exposes an api which can be used from the outside.

It exposes a playground ui at http://localhost:8080/ which can be used to test some queries:

Example requests:
__login__ (admin, admin is automatically available)
```
mutation {
  login(username: "admin", password: "admin") {
    id
    token
  }
}
```
returns  
```
{
  "data": {
    "login": {
      "id": "311a96cd-f01f-4148-8454-c7df62577f7d",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc0FkbWluIjp0cnVlLCJ1c2VySWQiOiIzMTFhOTZjZC1mMDFmLTQxNDgtODQ1NC1jN2RmNjI1NzdmN2QiLCJleHBpcmVzIjoxNjAzMTM5MTIyfQ.IMPTE6RmGDIrLX56UQ7-5luXCnWtkomlVr-s-P0DKsM"
    }
  }
}
```

Now you need to use this token to craft a bearer header.   
You can do this by setting the HTTP Headers at the bottom of the playground ui:
```
{
  "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc0FkbWluIjp0cnVlLCJ1c2VySWQiOiIzMTFhOTZjZC1mMDFmLTQxNDgtODQ1NC1jN2RmNjI1NzdmN2QiLCJleHBpcmVzIjoxNjAzMTM5MTIyfQ.IMPTE6RmGDIrLX56UQ7-5luXCnWtkomlVr-s-P0DKsM"
}
```

After that you can get a user or, if you are admin, create a new one.
__get user__
```
query {
  user(id:"311a96cd-f01f-4148-8454-c7df62577f7d"){
    id, name, email
  }
}
```

__create user__
```
mutation {
 	createUser(newUser:{
    username: "aligator",
    email: "aligator@suncraft-server.de",
    password: "superpassword",
    isAdmin: false
  }){
    id
  }
}
```