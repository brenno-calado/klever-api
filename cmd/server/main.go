package main

import (
	"context"
	"klever-api/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSendMessageServer
}

func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("received message: ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}

	return response, nil
}
func (service *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	Listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal(err)
	}

	grpc_error := grpcServer.Serve(Listener)
	if grpc_error != nil {
		log.Fatal(grpc_error)
	}
}