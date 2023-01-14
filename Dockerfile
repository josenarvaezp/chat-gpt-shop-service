FROM golang:latest

# Create app directory
RUN mkdir /app

# Add the source code
ADD . /app/

# Change working directory
WORKDIR /app

# Build the app (path to main.go not added by chatgpt)
RUN go build -o shop-service ./server/main.go

# Expose port
EXPOSE 9090

# Run the app
CMD ["./shop-service"]
