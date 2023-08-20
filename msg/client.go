package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "msg" // Protobuf paketinize göre düzenleyin
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCommunicationServiceClient(conn)
	request := &pb.Request{Message: "Hello, server!"}

	response, err := client.SendMessage(context.Background(), request)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Println("Server response:", response.Reply)
}
