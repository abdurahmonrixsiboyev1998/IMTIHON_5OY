package booking

import (
	"api_geteway/internal/protos/booking"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hotel() booking.BookHotelClient {
	conn, err := grpc.NewClient("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := booking.NewBookHotelClient(conn)
	return client
}
