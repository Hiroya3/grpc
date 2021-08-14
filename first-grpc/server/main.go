package main

import (
	"context"
	"errors"
	pb "github.com/Hiroya3/grpc/first-grpc/grpc-define"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedBornYearServer
}

func (b *server) GetBornYear(_ context.Context, age *pb.Age) (*pb.Year, error) {
	currentYear := time.Now().Year()
	returnYear := int32(currentYear) - age.GetAge()
	if returnYear < 0 {
		return nil, errors.New("you're not born")
	}
	return &pb.Year{Year: returnYear}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBornYearServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
