# microservice-structure

> An opinionated microservice reference implementation

Having spent some time developing microservices I have tested a lot of different ways to strucutre a microservice project.
With this project I want to gradually build a good and stable reference implementation for my future microservices.

#### Requirements
+ gRPC as main transport protocol
  + API defined by protobuf only
  + Provide a good set of default-interceptors build-in:
    + logging
    + payload logging
    + panic recovering
    + metrics (Prometheus)
    + tracing
    + caching
+ HTTP/JSON must be provided as well to support arbitrary clients
  + For easy integration into Apiary, I want to have a Swagger spec generated
+ **opinion**: API documentation does *not* live inside the service repository
  +  protobuf and swagger definitions should be stored in separate repositories.  
  I prefer to have the API specs of the whole application in one repository acting as single source of truth.
  When delegating work I can just write up the protobuf, push it to the repo and force the service to implement that contract.
  This is better because every developer gets notified about API changes inside the protobuf repository.
  If the protobufs are stored inside service/stack repositories you will most likely not be notified if the API changed.
  Also this enables me to version the API separately from the service. On top of that, you have central control of how the protobuf stubs are generated.
+ **opinion**: No frameworks
  + I've pretty much tried them all. A framework either forces you into a specific way of writing services or introduces
a lot of verbosity. Although I really like the concepts of `go-kit` i found it being too verbose - at least for now. Looking into the future I could imagine 
going back to go-kit. But currently I do not want any framework. 
+ MySQL and MongoDB integration
+ multistage Dockerfile
+ Jenkins is used for CI/CD
+ Services will run on Kubernetes

## Where are the protobufs / swagger specs?
As I've stated above, i keep them separate. You can find them here: [lukasjarosch/microservice-structure-protobuf](https://github.com/lukasjarosch/microservice-structure-protobuf)

That repository (contract-repository) defines the APIs of all our services.

## Features
 - [x] ENV only configuration
 - [x] gRPC server
   - [x] zap logging interceptor
   - [x] payload logging interceptor
   - [x] prometheus interceptor
 - [x] grpc-gateway integration
 - [x] grpc-gateway swagger setup
 - [x] graceful shutdown
 - [ ] Kubernetes deployment and service config
 - [ ] Jenkins pipeline script
 - [ ] extensive Makefile
 
 
