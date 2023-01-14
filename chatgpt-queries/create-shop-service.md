Query: write go code that implement the shop service above

Response:
```
package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type server struct{}

var bananas = []*Banana{
	&Banana{Id: "1", Color: "Yellow"},
	&Banana{Id: "2", Color: "Green"},
	&Banana{Id: "3", Color: "Yellow"},
}

func (s *server) ListBananas(ctx context.Context, req *ListBananasRequest) (*ListBananasResponse, error) {
	return &ListBananasResponse{Bananas: bananas[:req.Size]}, nil
}

func (s *server) GetBananaColor(ctx context.Context, req *GetBananaColorRequest) (*GetBananaColorResponse, error) {
	for _, banana := range bananas {
		if banana.Id == req.BananaId {
			return &GetBananaColorResponse{Color: banana.Color}, nil
		}
	}
	return nil, grpc.Errorf(grpc.Code(NotFound), "Banana not found")
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterShopServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

This is an example implementation of the "Shop" service in Go using gRPC. The server struct implements the two methods defined in the proto file: "ListBananas" and "GetBananaColor".
The ListBananas method takes the size from the request and returns the slice of bananas upto that size.
The GetBananaColor method takes the banana_id from the request and iterates over the bananas to find the matching banana, and if found, it returns the color of that banana, otherwise it returns an error message "Banana not found".
You'll have to import and use the protobuf generated go files in the imports and structs and methods.

Please note that I'm returning hardcoded data here, In a real-world scenario you'll probably want to replace this with data from a database or some other data store.