package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "msg" 
)

type server struct{}

func (s *server) SendMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	message := req.GetMessage()
	response := &pb.Response{Reply: fmt.Sprintf("You said: %s", message)}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterCommunicationServiceServer(s, &server{})
	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
