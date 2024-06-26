package main

import (
	"log"
	"context"
	"time"
	pb "github.com/Kevincxv/Go-gRPC-Demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Unary Started")
	log.Printf("%s", res.Message)
	log.Printf("Unary Finished")
}