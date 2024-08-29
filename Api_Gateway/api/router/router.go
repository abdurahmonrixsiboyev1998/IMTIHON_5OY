package router

import (
	"api_geteway/config"
	"api_geteway/internal/connections"
	jwttoken "api_geteway/internal/jwt"
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter() {
	c := config.LoadConfig()

	r := http.NewServeMux()
	handler := connections.NewHandler()

	// Users
	r.Handle("/users/register", http.HandlerFunc(handler.Register))
	r.Handle("/users/verify", http.HandlerFunc(handler.Verify))
	r.Handle("/users/login", http.HandlerFunc(handler.LogIn))
	r.Handle("/users/", jwttoken.JWTMiddleware(http.HandlerFunc(handler.GetUser)))
	r.Handle("/users/update", jwttoken.JWTMiddleware(http.HandlerFunc(handler.UpdateUser)))
	r.Handle("/users/delete", jwttoken.JWTMiddleware(http.HandlerFunc(handler.DeleteUser)))
	r.Handle("/users/logout", jwttoken.JWTMiddleware(http.HandlerFunc(handler.LogOut)))

	// Swagger
	r.Handle("/swagger/", httpSwagger.WrapHandler)

	// Hotel
	r.Handle("/hotels/create", jwttoken.JWTMiddleware(http.HandlerFunc(handler.CreateHotel)))
	r.Handle("/hotels/rooms/create", jwttoken.JWTMiddleware(http.HandlerFunc(handler.CreateRoom)))
	r.Handle("/hotels", jwttoken.JWTMiddleware(http.HandlerFunc(handler.GetHotels)))
	r.Handle("/hotels/", jwttoken.JWTMiddleware(http.HandlerFunc(handler.GetHotel)))
	r.Handle("/hotels/rooms/", jwttoken.JWTMiddleware(http.HandlerFunc(handler.GetRooms)))
	r.Handle("/hotels/room", jwttoken.JWTMiddleware(http.HandlerFunc(handler.GetRoom)))
	r.Handle("/hotels/update", jwttoken.JWTMiddleware(http.HandlerFunc(handler.UpdateHotel)))
	r.Handle("/hotels/rooms/update", jwttoken.JWTMiddleware(http.HandlerFunc(handler.UpdateRoom)))
	r.Handle("/hotels/delete", jwttoken.JWTMiddleware(http.HandlerFunc(handler.DeleteHotel)))
	r.Handle("/hotels/rooms/delete", jwttoken.JWTMiddleware(http.HandlerFunc(handler.DeleteRoom)))

	// Booking
	r.Handle("/bookings", jwttoken.JWTMiddleware(http.HandlerFunc(handler.CreateBooking)))
	r.Handle("/waitinglists", jwttoken.JWTMiddleware(http.HandlerFunc(handler.CreateWaiting)))
	r.Handle("/bookings/", jwttoken.JWTMiddleware(http.HandlerFunc(handler.GetBooking)))
	r.Handle("/waitinglists/", jwttoken.JWTMiddleware(http.HandlerFunc(handler.GetWaiting)))
	r.Handle("/bookings/update", jwttoken.JWTMiddleware(http.HandlerFunc(handler.UpdateBooking)))
	r.Handle("/waitinglists/update", jwttoken.JWTMiddleware(http.HandlerFunc(handler.UpdateWaiting)))
	r.Handle("/bookings/delete", jwttoken.JWTMiddleware(http.HandlerFunc(handler.DeleteBooking)))
	r.Handle("/waitinglists/delete", jwttoken.JWTMiddleware(http.HandlerFunc(handler.DeleteWaiting)))

	fmt.Printf("Server started on port %s\n", c.User.Port)
	if err := http.ListenAndServe(c.User.Port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
