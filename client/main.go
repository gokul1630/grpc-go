package main

import (
	"context"
	"gorpc/gorpc"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Printf("failed to listen: %s", err)
		return
	}

	defer conn.Close()
	client := gorpc.NewMessageClient(conn)

	res, err := client.SayHello(context.TODO(), &gorpc.Request{Name: "goku"})
	if err != nil {
		log.Printf("failed to listen: %s", err)
		return
	}

	// result
	log.Println(res.GetName())

}
