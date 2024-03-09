package main

import (
	"log"
	"time"

	pb "github.com/zigamedved/go-grpc-project/proto"
)

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
