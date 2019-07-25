package main

import (
	"grpcTest/configServer/dao"
	"grpcTest/rpgcServer/dbConfigInterface"
	"log"
	"net"
	_ "grpcTest/rpgcServer/dbConfigInterface"
	_ "grpcTest/configServer/dao"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	dbConfigInterface.RegisterDataSourceServiceServer(s, &dao.Server{})


	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
