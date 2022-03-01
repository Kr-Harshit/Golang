package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("[INFO] Recieve Message Body from client: %s", in.Body)
	return &Message{Body: "Hello From The Server!"}, nil
}
