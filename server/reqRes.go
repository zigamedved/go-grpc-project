package main

import (
	"context"

	pb "github.com/zigamedved/go-grpc-project/proto"
)

func (s *helloServer) RequestResponse(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, world!"}, nil
}
