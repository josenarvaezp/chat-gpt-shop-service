Query: write a cli written in go that can communicate with the service above. The cli should have a command called listBananas to get the list of bananas and a command getBananaColor to get the color of a specific banana

Response:
```
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/your_package_name/your_proto_package_name"
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

	client := your_proto_package_name.NewShopClient(conn)

	if size > 0 {
		listBananas(client)
	} else if bananaId != "" {
		getBananaColor(client)
	} else {
		log.Fatalf("No arguments provided")
	}
}

func listBananas(client your_proto_package_name.ShopClient) {
	response, err := client.ListBananas(context.Background(), &your_proto_package_name.ListBananasRequest{Size: size})
	if err != nil {
		log.Fatalf("Error while calling ListBananas: %v", err)
	}
	for _, banana := range response.Bananas {
		fmt.Printf("Banana %s is %s\n", banana.Id, banana.Color)
	}
}

func getBananaColor(client your_proto_package_name.ShopClient) {
	response, err := client.GetBananaColor(context.Background(), &your_proto_package_name.GetBananaColorRequest{BananaId: bananaId})
	if err != nil {
		log.Fatalf("Error while calling GetBananaColor: %v", err)
	}
	fmt.Printf("Banana %s is %s\n", bananaId, response.Color)
}
```

This is an example CLI written in Go that can communicate with the "Shop" service. The CLI uses the gRPC library to connect to the server and call the methods. The CLI takes command-line arguments to specify the banana ID and the number of bananas to retrieve. The address of the server can also be specified using the address flag and defaults to "localhost:9090".
You'll have to import and use the protobuf generated go files in the imports and structs and methods.
The listBananas function takes the size from the command-line argument and calls the ListBananas method on the server, then prints the result.
The getBananaColor function takes the bananaId from the command-line argument and calls the GetBananaColor method on the server