package connections

import (
	broadcas "api_geteway/internal/broadcast"
	"api_geteway/internal/client/booking"
	"api_geteway/api/handler"
	hotel_service "api_geteway/internal/client/hotel"
	user_service "api_geteway/internal/client/user"
	redis_method "api_geteway/internal/redis/method"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewBroadcast() *broadcas.Adjust {
	u := user_service.UserClinet()
	h := hotel_service.Hotel()
	b := booking.Hotel()
	r := Redis()
	ctx := context.Background()
	return &broadcas.Adjust{U: u, Ctx: ctx, R: r, H: h, B: b}
}

func NewHandler() *handler.Handler {
	a := NewBroadcast()
	return &handler.Handler{B: a}
}

func Redis() *redis_method.Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	return &redis_method.Redis{R: client, Ctx: ctx}
}
