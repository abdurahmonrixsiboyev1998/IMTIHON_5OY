package services

import (
	"context"
	"fmt"
	"log"
	"notify_service/api/handler"
	email1 "notify_service/internal/email"
	"notify_service/internal/kafka/producer"
	"notify_service/internal/protos/notification"
	"strconv"
)

type Service struct {
	notification.UnimplementedNotificationServer
	P *handler.WebSocket
}

func (s *Service) AddUser(ctx context.Context, req *notification.AddnewUser) (*notification.EMailSendResponse, error) {
	fmt.Println("there is request")
	if err := s.P.AddUser(req.UserId, nil); err != nil {
		log.Println(err)
		return nil, err
	}
	return &notification.EMailSendResponse{
		Message: "User added successfully",
	}, nil

}

func (u *Service) Notification(ctx context.Context, req *notification.ProduceMessage) (*notification.EMailSendResponse, error) {
	if err := producer.Producer(strconv.Itoa(int(req.UserId)), req.Message); err != nil {
		log.Println(err)
		return nil, err
	}

	return &notification.EMailSendResponse{Message: "succesfully"}, nil
}

func (u *Service) Email(ctx context.Context, req *notification.EmailSend) (*notification.EMailSendResponse, error) {
	fmt.Println("request is came")
	if err := email1.Sent(req.Email, req.Message); err != nil {
		log.Println(err)
		return nil, err
	}
	
	return &notification.EMailSendResponse{Message: "successfully"}, nil
}
