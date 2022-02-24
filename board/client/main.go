package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "example.com/othello/board"
)

const (
	defaultName = "world"
	addr = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBoardApiClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Reverse(ctx, &pb.Board{
		Stone: "b",
		X: 3,
		Y: 5,
		Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.GetSquares())
	r, err = c.Putable(ctx, &pb.Board{
		Stone: "b",
		Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.GetSquares())
}