package main

import (
	pb "grpcTest/rpgcServer/dbConfigInterface"
	"google.golang.org/grpc"
	"context"
	"log"
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

	param := &pb.DataSourceRequest{}

	param.Id= "10"
	param.DataName="20"
	param.JdbcUrl="30"
	param.DriverClass="40"
	param.Password="50"
	param.Remake="60"
	param.User="70"
	param.Status=80
	param.CreateTime=1231230
	param.WriteOrRead=0

	c := pb.NewDataSourceServiceClient(conn)
	r, err := c.FindAllDataSource(context.Background(), param)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}