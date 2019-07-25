package dao

import (
	pb "grpcTest/rpgcServer/dbConfigInterface"
	"log"

	"context"
)

func (s *Server) FindAllDataSource(ctx context.Context, in *pb.DataSourceRequest) (*pb.DataSourceReply, error) {

	requuest := &pb.DataSourceReply{}
	status := &pb.Status{}
	rd := &pb.DataSourceRequest{}

	log.Printf("Greeting: %s", in)

	requuest.SInfo= status
	requuest.DrInfo = append(requuest.DrInfo, rd)

	status.ResCode = "1"
	status.ResContent="2"
	status.ResMsg="3"

	rd.Id= "1"
	rd.DataName="2"
	rd.JdbcUrl="3"
	rd.DriverClass="4"
	rd.Password="5"
	rd.Remake="6"
	rd.User="7"
	rd.Status=8
	rd.CreateTime=123123
	rd.WriteOrRead=1


	return requuest, nil
}


func (s *Server) AddDataSource(context.Context, *pb.DataSourceRequest) (*pb.DataSourceReply, error) {
	panic("implement me")
}

func (s *Server) UpdateDataSource(context.Context, *pb.DataSourceRequest) (*pb.DataSourceReply, error) {
	panic("implement me")
}

func (s *Server) DelDataSource(context.Context, *pb.DataSourceRequest) (*pb.DataSourceReply, error) {
	panic("implement me")
}

func (s *Server) CountDataSource(context.Context, *pb.DataSourceRequest) (*pb.DataSourceReply, error) {
	panic("implement me")
}
