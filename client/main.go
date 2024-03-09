package main

import (
	"log"

	pb "github.com/zigamedved/go-grpc-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	client := pb.NewServiceClient(conn)
	names := &pb.NamesList{Names: []string{"Joe", "Doe"}}
	//CallRequestResponse(client)
	//CallServerSideStreaming(client, names)
	CallClientSideStreaming(client, names)
}
