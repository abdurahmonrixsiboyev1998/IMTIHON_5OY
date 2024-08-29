package notification17

import (
	notification "api_geteway/internal/protos/notification"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hotel() notification.NotificationClient {
	conn, err := grpc.NewClient("localhost:8877", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("notification error",err)
	}
	client := notification.NewNotificationClient(conn)
	return client
}
