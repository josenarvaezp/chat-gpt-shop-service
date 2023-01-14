# Shop Service

Most of this service (including this readme) was writen by chatGPT using this conversation: [chatGPT conversation](https://github.com/josenarvaezp/chat-gpt-shop-service/blob/master/chatgpt-conversation.md). The changes made to chatGPT's output can be found here: [changes log](https://github.com/josenarvaezp/chat-gpt-shop-service/blob/master/changes-log.md).

This is the README file for the Shop service. The Shop service is a gRPC service that allows clients to list and get the color of bananas.

## Running the service using Docker
To run the Shop service using Docker, you will first need to have Docker installed on your machine.

Build the Docker image for the Shop service by running the following command in the root directory of the project:
```
docker build -t shop-service .
```

Run the Docker image by running the following command:
```
docker run -p 9090:9090 shop-service
```

This will start the Shop service on port 9090.

## Interacting with the service using the CLI
To interact with the Shop service, you will need to have the CLI for the service installed on your machine.

## Build CLI (not added by chatgpt)
To build CLI run:
```
go build -o ./shop-cli ./cli/main.go
```

### Listing bananas
To list the bananas, you can run the following command:
(chatgpt command was ```./shop-cli listBananas-size=<number of bananas>``` but doesn't work)
```
./shop-cli -size=<number of bananas>
```
Replace <number of bananas> with the number of bananas you want to retrieve.

### Getting the color of a banana
To get the color of a banana, you can run the following command:
(chatgpt command was ```./shop-cli getBananaColor -bananaId=<banana id>``` but doesn't work)
```
./shop-cli getBananaColor -bananaId=<banana id>
```
Replace <banana id> with the ID of the banana you want to retrieve the color of.

You can also check the help command of the cli by running the command ./shop-cli -h