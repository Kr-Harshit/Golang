package main

import (
	"context"
	"log"

	"github.com/Kr_Harshit/golang-example/grpc/chat"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[ERROR] did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := &chat.Message{
		Body: "hello from client!",
	}

	response, err := c.SayHello(context.Background(), message)
	if err != nil {
		log.Fatalf("[ERROR] calling SayHello failed!\n")
	}

	log.Printf("[INFO] Response from server %s", response.Body)
}
