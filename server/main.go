package main

import (
	"context"
	"fmt"
	"gorpc/gorpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	gorpc.UnimplementedMessageServer
}

func (s *server) SayHello(ctx context.Context, req *gorpc.Request) (*gorpc.Response, error) {
	fmt.Println(ctx.Deadline())
	return &gorpc.Response{Name: fmt.Sprintf("Hello %s!", req.GetName())}, nil
}

func main() {
	s := grpc.NewServer()
	gorpc.RegisterMessageServer(s, &server{})

	listen, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Printf("failed to listen: %s", err)
		return
	}

	s.Serve(listen)
}
