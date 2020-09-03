package main

import (
	"fmt"
	"log"
	"net"

	"../calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:7749")
	if err != nil {
		log.Fatalf("err", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	fmt.Println("Server running")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err", err)
	}
}
