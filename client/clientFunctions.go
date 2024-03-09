package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/zigamedved/go-grpc-project/proto"
)

func CallRequestResponse(client pb.ServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.RequestResponse(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("error while calling request response: %v", err)
	}

	log.Printf("%s", res.Message)
}

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

func CallClientSideStreaming(client pb.ServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started...")
	stream, err := client.ClientSideStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("Sent request")
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while sending: %v", err)
	}
	log.Printf("...client streaming ended")
	log.Printf("%v", res.Messages)
}
