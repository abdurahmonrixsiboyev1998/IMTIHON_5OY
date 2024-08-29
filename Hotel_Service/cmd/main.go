package main

import (
	"fmt"
	"hotel_service/config"
	"hotel_service/internal/connections"
	"hotel_service/internal/protos/hotel"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.LoadConfig()

	ls, err := net.Listen(c.User.Host, c.User.Port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	server := connections.NewGrpc()

	hotel.RegisterHotelServer(s,server)
	reflection.Register(s)
	
	fmt.Printf("server started on the port %s", c.User.Port)

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
