package main

import (
	"context"
	"log"
	"net"

	"github.com/josenarvaezp/chat-gpt-shop-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct {
	proto.UnimplementedShopServer
}

var bananas = []*proto.Banana{
	{Id: "1", Color: "Yellow"},
	{Id: "2", Color: "Green"},
	{Id: "3", Color: "Yellow"},
}

func (s *server) ListBananas(ctx context.Context, req *proto.ListBananasRequest) (*proto.ListBananasResponse, error) {
	return &proto.ListBananasResponse{Bananas: bananas[:req.Size]}, nil
}

func (s *server) GetBananaColor(ctx context.Context, req *proto.GetBananaColorRequest) (*proto.GetBananaColorResponse, error) {
	for _, banana := range bananas {
		if banana.Id == req.BananaId {
			return &proto.GetBananaColorResponse{Color: banana.Color}, nil
		}
	}

	// code not correctly added by chatgpt
	return nil, grpc.Errorf(codes.NotFound, "Banana not found")
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterShopServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
