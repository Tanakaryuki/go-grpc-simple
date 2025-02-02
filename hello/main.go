package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/Tanakaryuki/go-grpc-simple/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello World"}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, &server{})
	log.Printf("Hello service listening on :%s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
