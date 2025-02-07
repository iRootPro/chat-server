package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/irootpro/chat-server/pkg/servers/grpc/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 3009

type server struct {
	desc.UnimplementedChat_V1Server
}

func (s server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("usernames: %v\n", req.Usernames)
	return &desc.CreateResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChat_V1Server(s, server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server %v", err)
	}
}
