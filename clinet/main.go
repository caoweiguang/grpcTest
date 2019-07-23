package main

import (
	pb "ch2/helloworld"
	"context"
	"log"

	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {
	//建立一个连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "123"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
