package grpchandle

import (
	"app/generated/grpc/servicegrpc"
	"app/internal/connection"
)

type grpcHandle struct {
	servicegrpc.UnimplementedStreamServiceServer
	infoConnection connection.Connection
}

func Register() servicegrpc.StreamServiceServer {
	return &grpcHandle{
		infoConnection: connection.GetConnect(),
	}
}
