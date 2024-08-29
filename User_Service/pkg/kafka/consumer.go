package kafka

import (
	"context"
	"encoding/json"
	"log"
	"user_service/internal/models"
	"user_service/pkg/protos/user"

	grpcmethods "user_service/pkg/methods"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer struct {
	C   *grpcmethods.Service
	Ctx context.Context
}

func (u *Consumer) Consumer() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics("hoteluser17"),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	for {
		fetches := client.PollFetches(ctx)
		if err := fetches.Errors(); len(err) > 0 {
			log.Fatal(err)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			for _, record := range ftp.Records {
				if err := u.Adjust(record); err != nil {
					log.Println(err)
				}
			}
		})
	}
}

func (u *Consumer) Adjust(record *kgo.Record) error {
	switch string(record.Key) {
	case "create":
		if err := u.Create(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "update":
		if err := u.Update(record.Value); err != nil {
			log.Println(err)
			return err
		}
	case "delete":
		if err := u.Delete(record.Value); err != nil {
			log.Println(err)
			return err
		}
	case "logout":
		if err := u.LogOut(record.Value); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (u *Consumer) Create(req []byte) error {
	var req1 models.RegisterUserRequest

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = user.RegisterUserRequest{
		Username: req1.Username,
		Age:      req1.Age,
		Email:    req1.Email,
		Password: req1.Password,
	}
	_, err := u.C.Register(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer) Update(req []byte) error {
	var req1 models.UpdateUserRequest

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = user.UpdateUserRequest{
		Id:       req1.ID,
		Username: req1.Username,
		Age:      req1.Age,
		Email:    req1.Email,
		Password: req1.Password,
	}
	_, err := u.C.UpdateUser(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer) LogOut(req []byte) error {
	var req1 models.GetUserRequest
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = user.GetUserRequest{
		Id: req1.ID,
	}

	_, err := u.C.LogOut(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer) Delete(req []byte) error {
	var req1 models.GetUserRequest
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = user.GetUserRequest{
		Id: req1.ID,
	}

	_, err := u.C.DeleteUser(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
