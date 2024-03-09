package main

import (
	"context"
	"io"
	"log"

	pb "github.com/zigamedved/go-grpc-project/proto"
)

func CallServerSideStreaming(client pb.ServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started...")
	stream, err := client.ServerSideStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}
		log.Println(message)
	}

	log.Printf("...streaming ended")
}
