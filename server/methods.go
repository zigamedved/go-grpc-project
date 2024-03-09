package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/zigamedved/go-grpc-project/proto"
)

func (s *helloServer) RequestResponse(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, world!"}, nil
}

func (s *helloServer) ServerSideStreaming(req *pb.NamesList, stream pb.Service_ServerSideStreamingServer) error {
	log.Printf("got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hey " + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(3 * time.Second)
	}
	return nil
}

func (s *helloServer) ClientSideStreaming(stream pb.Service_ClientSideStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}

		log.Printf("Got request with name: %v", req.Name)
		messages = append(messages, req.Name)
	}
}

func (s *helloServer) BidirectionalStreaming(stream pb.Service_BidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hey " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
