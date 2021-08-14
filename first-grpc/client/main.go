package main

import (
	"context"
	pb "github.com/Hiroya3/grpc/first-grpc/grpc-define"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	address    = "localhost:50051"
	defaultAge = 20
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewBornYearClient(conn)

	age := defaultAge
	if len(os.Args) > 1 {
		age, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Args's type was invalid: %v", err)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetBornYear(ctx, &pb.Age{Age: int32(age)})
	if err != nil {
		log.Fatalf("could not get born year, %v", err)
	}
	log.Printf("you were born in %v", r.GetYear())
}
