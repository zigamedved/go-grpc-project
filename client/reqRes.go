package main

import (
	"context"
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
