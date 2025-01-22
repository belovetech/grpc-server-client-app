package main

import (
	"context"
	"log"
	"net"
	pb "prac-three/gen/hello"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
	listneer, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpServer, &server{})

	log.Println("Server started on port :50051")
	if err := grpServer.Serve(listneer); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
