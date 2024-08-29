package connections

import (
	"booking_service/config"
	user_service "booking_service/internal/clients/user"
	hotel_service "booking_service/internal/clients/hotel"
	notify "booking_service/internal/clients/notification"
	"booking_service/internal/database/method"
	consumer "booking_service/pkg/brokers/consumer"

	methods "booking_service/internal/booking_methods"
	interfac "booking_service/internal/interface"
	interfaceservices "booking_service/internal/services"
	"booking_service/pkg/server"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabase() interfac.Booking {
	c := config.LoadConfig()
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.DBname))
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	return &method.Database{Db: db}
}

func NewService() *interfaceservices.Database {
	a := NewDatabase()
	return &interfaceservices.Database{D: a}
}

func NewAdjust() interfac.BookingAdjust {
	a := NewService()
	user := user_service.UserClinet()
	hotel := hotel_service.Hotel()
	n := notify.Hotel()
	return &server.Adjust{
		S: a, User: user,
		Hotel: hotel, N: n,
	}
}

func NewAdjus() *interfaceservices.AdjustDatabase {
	a := NewAdjust()
	return &interfaceservices.AdjustDatabase{A: a}
}

func NewGrpc() *methods.Grpc {
	a := NewAdjus()
	return &methods.Grpc{A: a}
}

func NewConsumer() *consumer.Consumer17 {
	a := NewAdjus()
	ctx := context.Background()
	return &consumer.Consumer17{A: a, Ctx: ctx}
}
