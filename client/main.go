package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/scotty-c/abstract/proto"

	"google.golang.org/grpc"
)

func readJSONFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	// Replace 'localhost:50051' with the address of your gRPC server
	serverAddress := "localhost:50051"
	jsonFilePath := "task_def.json" // Change this to your JSON file path

	// Read JSON data from file
	jsonData, err := readJSONFile(jsonFilePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Set up a connection to the server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewNetworkClient(conn)

	// Create a request
	request := &pb.JsonRequest{
		JsonData: jsonData,
	}

	// Call the gRPC service
	response, err := client.SendJsonData(context.Background(), request)
	if err != nil {
		log.Fatalf("Error calling gRPC service: %v", err)
	}

	fmt.Printf("Response received: %s\n", response.ResponseMessage)
}
