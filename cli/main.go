package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/josenarvaezp/chat-gpt-shop-service/proto"
)

var (
	bananaId string
	size     int
	address  string
)

func init() {
	flag.StringVar(&bananaId, "bananaId", "", "The ID of the banana to get the color of")
	flag.IntVar(&size, "size", -1, "The number of bananas to get")
	flag.StringVar(&address, "address", "localhost:9090", "The address of the server")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewShopClient(conn)

	if size > 0 {
		listBananas(client)
	} else if bananaId != "" {
		getBananaColor(client)
	} else {
		log.Fatalf("No arguments provided")
	}
}

func listBananas(client proto.ShopClient) {
	// int32 cast not added by chatgpt
	response, err := client.ListBananas(context.Background(), &proto.ListBananasRequest{Size: int32(size)})
	if err != nil {
		log.Fatalf("Error while calling ListBananas: %v", err)
	}
	for _, banana := range response.Bananas {
		fmt.Printf("Banana %s is %s\n", banana.Id, banana.Color)
	}
}

func getBananaColor(client proto.ShopClient) {
	response, err := client.GetBananaColor(context.Background(), &proto.GetBananaColorRequest{BananaId: bananaId})
	if err != nil {
		log.Fatalf("Error while calling GetBananaColor: %v", err)
	}
	fmt.Printf("Banana %s is %s\n", bananaId, response.Color)
}
