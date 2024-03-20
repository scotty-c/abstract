# Abstract 

This is a POC and nothing more. It is not intended to be used in production and is not a complete solution. It is a proof of concept to showcase server side abstractions and what we can build ontop of what exsisits in AWS.

## Introduction
Abstract is a proof of concept abstraction over AWS networking with the focus on ECS. It takes a standard ECS [task definition](client/task_def.json) and creates the necessary networking infrastructure to support it. This includes VPC, subnets, security groups, and an ALB.
Abstract is broken into two main components: the client and the server. The client has very little logic and reads the task definition from a file and passes it to the server via grpc. The server is responsible for creating the networking infrastructure and is written in Go.
The point of this POC is showcase server side abstractions and what we can build ontop of what exsisits in AWS. The server is designed to be a generic abstraction over AWS networking and can be extended to support other AWS services besides ECS.

## Design
As this is a moc of a larger system, the design is very simple. The client reads the task definition from a file and passes it to the server via grpc. The server then creates the necessary networking infrastructure to support the task definition. From the task definition, we can make assumptions about the networking infrastructure that is needed. For example, we can assume that the task definition will need a VPC, subnets, security groups, ports to expose etc. As this POC and runs locally you will need to have AWS credentials set up on your machine, in a production this would not be true. 

## Running the POC
To run the POC, you will need to have the following installed:
- Go
- AWS credentials

To run the server:
```bash
cd server && go run main.go
```
In a new terminal, run the client:
```bash
cd client && go run main.go
```