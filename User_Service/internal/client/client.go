package client

import (
	"log"
	notification "user_service/pkg/protos/notification"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Hotel() notification.NotificationClient {
	c, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("notification error",err)
	}
	client := notification.NewNotificationClient(c)
	return client
}
