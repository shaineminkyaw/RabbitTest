package rpc

import (
	"context"

	"github.com/shaineminkyaw/RabbitTest/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *GrpcServer) Register(ctx context.Context, req *pb.RequestUser) (*pb.ResponseUser, error) {
	//

	email := req.GetEmail()
	password := req.GetPassword()

	if email == "" {
		return nil, status.Errorf(codes.Internal, "no found data from request.....")
	}

	if password == "" {
		return nil, status.Errorf(codes.Internal, "no found password from request ...")
	}

	return &pb.ResponseUser{
		Status: "Ok I got user information......",
	}, nil
}
