# daily-shit
DailyShit is an example microservice newspaper.

* frontend:
    * website
* gateway
* microservices
    * login service
    * post service
    * comment service
    * newsletter service
    * ...
    
website <-> Gateway (graphql)  
gateway <-> microservices (gRPC)  
microservice <-> microservice (nats) 