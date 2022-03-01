package main

import (
	"log"
	"net"

	"github.com/Kr_Harshit/golang-example/grpc/chat"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("[ERROR] Failed  to listen %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()

	// Registring Services
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[ERROR] Failed to serve: %v", err)
	}

	log.Printf("[INFO] gRPC server running at %s", port)
}
