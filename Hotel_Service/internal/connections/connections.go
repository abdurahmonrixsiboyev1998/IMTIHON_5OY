package connections

import (
	"database/sql"
	"fmt"
	"hotel_service/config"
	"hotel_service/internal/databases/methods"
	interfac "hotel_service/internal/interface"
	adjustservice "hotel_service/internal/server"
	grpcmethod "hotel_service/internal/method"
	"hotel_service/internal/services"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabase() interfac.Hotel {
	c := config.LoadConfig()
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.DBname))
	if err != nil {
		log.Println(err)
	}

	if err := db.Ping(); err != nil {
		log.Println(err)
	}

	return &methods.Database{Db: db}
}

func NewService() *services.Database {
	a := NewDatabase()
	return &services.Database{S: a}
}

func NewAdjust() interfac.Adjust {
	a := NewService()
	return &adjustservice.Adust{S: a}
}

func NewAdjustService() *services.Adjust {
	a := NewAdjust()
	return &services.Adjust{A: a}
}

func NewGrpc() *grpcmethod.GrpcService {
	a := NewAdjustService()
	return &grpcmethod.GrpcService{A: a}
}
