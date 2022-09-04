package grpc

import (
	"log"
	"net"
	"user/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	pb.UnimplementedUserRegisterServer
}

func NewGRPCServer() (*GrpcServer, error) {
	return &GrpcServer{}, nil
}

func RunGRPC() {
	resource, err := NewGRPCServer()
	if err != nil {
		log.Println(err)
	}

	server := grpc.NewServer()
	pb.RegisterUserRegisterServer(server, resource)
	reflection.Register(server)

	lis, err := net.Listen("tcp", "localhost:7070")
	if err != nil {
		log.Println(err)
	}

	log.Println("Starting grpc server localhost:7070")
	err = server.Serve(lis)
	if err != nil {
		log.Println(err)
	}
}
