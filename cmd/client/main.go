package main

import (
	"context"
	"klever-api/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, client_err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if client_err != nil {
		log.Fatal(client_err)

	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello grpc",
	}

	res, request_err := client.RequestMessage(context.Background(), req)

	if request_err != nil {
		log.Fatal(request_err)
	}

	log.Print(res.GetStatus())
}
