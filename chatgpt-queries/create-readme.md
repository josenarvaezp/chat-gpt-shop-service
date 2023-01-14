Query: write a README file for the shop service explaining how to run the service using Docker and how to interact with it with the CLI

Response:
Shop Service
This is the README file for the Shop service. The Shop service is a gRPC service that allows clients to list and get the color of bananas.

Running the service using Docker
To run the Shop service using Docker, you will first need to have Docker installed on your machine.

Build the Docker image for the Shop service by running the following command in the root directory of the project:
Copy code
docker build -t shop-service .
Run the Docker image by running the following command:
Copy code
docker run -p 9090:9090 shop-service
This will start the Shop service on port 9090.

Interacting with the service using the CLI
To interact with the Shop service, you will need to have the CLI for the service installed on your machine.

Listing bananas
To list the bananas, you can run the following command:

Copy code
./shop-cli listBananas -size=<number of bananas>
Replace <number of bananas> with the number of bananas you want to retrieve.

Getting the color of a banana
To get the color of a banana, you can run the following command:

Copy code
./shop-cli getBananaColor -bananaId=<banana id>
Replace <banana id> with the ID of the banana you want to retrieve the color of.

You can also check the help command of the cli by running the command ./shop-cli -h

Please note that this is just an example, you may have different ways of running the service and interacting with it, you should adjust the commands and instructions according to your specific implementation.