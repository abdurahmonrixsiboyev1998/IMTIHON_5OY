package hotel_service

import (
	"booking_service/pkg/protos/hotel"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hotel() hotel.HotelClient {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := hotel.NewHotelClient(conn)
	return client
}
