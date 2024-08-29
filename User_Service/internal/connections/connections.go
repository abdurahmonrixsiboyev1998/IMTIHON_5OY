package connections

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	client "user_service/internal/client"
	"user_service/config"

	"user_service/internal/database/db"
	"user_service/internal/service"

	interfac "user_service/internal/interface"
	kafka "user_service/pkg/kafka"
	server "user_service/pkg/server"
	methods "user_service/pkg/methods"

	_ "github.com/lib/pq"
)

func NewDatabase() interfac.User {
	c := config.LoadConfig()
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.DBname))
	if err != nil {
		log.Println("database error", err)
	}

	if err := db.Ping(); err != nil {
		log.Println("ping error", err)
	}

	n := client.Hotel()
	return &post.Database{
		Db: db, N: n,
	}
}

func NewService() *service.Service {
	a := NewDatabase()
	return &service.Service{D: a}
}

func NewAdjust() interfac.Adjust {
	a := NewService()
	return &server.Adjust{S: a}
}

func NewAdjustService() *service.Adjust {
	a := NewAdjust()
	return &service.Adjust{A: a}
}

func NewGrpc() *methods.Service {
	a := NewAdjustService()
	return &methods.Service{S: a}
}

func NewConsumer() *kafka.Consumer {
	a := NewGrpc()
	ctx := context.Background()
	return &kafka.Consumer{
		C: a, Ctx: ctx,
	}
}
