# Log of changes to fix chatGPT output

- added markdown syntax to README
- Added to proto:
```
option go_package = "github.com/josenarvaezp/chat-gpt-shop-service/proto";
```
- added proto.UnimplementedShopServer to grpc server struct
- ran proto generation as 
```
protoc --go_out=. --go_opt=paths=source_relative \           
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./proto/shop.proto
```
- updated import paths
- Updated path to main.go in Dockerfile
- int32 cast in CLI not added by chatgpt
- added instructions on how to build cli in readme
- Update cli commands in README