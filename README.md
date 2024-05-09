# Go gRPC Demo


This project demonstrates the implementation of four types of communication between client and server using Go:

1. Unary Communication: Standard Request and Response architecture.
2. Server Streaming: The client requests and receives a stream of responses from the server.
3. Client Streaming: The client sends a stream of messages and receives a response from the server.
4. Bi-directional Streaming: Both the client and server simultaneously send streams of requests and responses to each other.

## Project Structure

- **/client**: Contains the source code for the gRPC client.
- **/server**: Contains the source code for the gRPC server.
- **/proto**: Contains the protocol buffer files, defining the service and message types used by the gRPC client and server.

## Prerequisites

- Go (1.14 or later recommended)
- Protocol buffer compiler (protoc)

## Installation

1. Clone the repository:
   git clone https://github.com/Kevincxv/Go-gRPC-Demo
   cd Go-gRPC-Demo
2. Install Dependencies:
   go mod tidy

## Compolling Protocol Buffers
- protoc --go_out=. --go-grpc_out=. proto/greet.proto

## Running The Server

- cd server
- go run server.go

## Running The CLient

- cd client
- go run client.go