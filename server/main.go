package main

import (
	"log"
	"net"

	pb "github.com/scotty-c/abstract/proto"
	"google.golang.org/grpc"

	"github.com/scotty-c/abstract/server/server" // Import the server package
)

func main() {
	// Set up a listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Create a new Server
	server := &server.Server{} // Change to a pointer to server.Server

	// Register the YourService server
	pb.RegisterNetworkServer(s, server)

	// Start the server
	log.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
