package main

import (
	"context"
	"errors"
	pb "github.com/Hiroya3/grpc/first-grpc/grpc-define"
	"time"
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

}
