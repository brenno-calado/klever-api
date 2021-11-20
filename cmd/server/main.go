package main

import (
	"context"
	"klever-api/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedExchangeServiceServer
}

func (service *Server) Exchange(ctx context.Context, req *pb.ExchangeRequest) (*pb.ExchangeResponse, error) {
	log.Println("Requested value: ", req.GetValue())
	log.Println("Requested Coin: ", req.GetCoinName())

	response := &pb.ExchangeResponse{
		Balance: 4500,
	}

	return response, nil
}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterExchangeServiceServer (grpcServer, &Server{})

	Listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal(err)
	}

	grpc_error := grpcServer.Serve(Listener)
	if grpc_error != nil {
		log.Fatal(grpc_error)
	}
}