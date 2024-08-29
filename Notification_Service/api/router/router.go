package router

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"notify_service/config"
	"notify_service/internal/connections"
	"notify_service/internal/protos/notification"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewRouter() {
	r := mux.NewRouter()
	a := connections.NewService().P

	r.HandleFunc("/ws", a.HandleWebSocket)
	go Grpc()

	fmt.Println("server started on port 7788")
	log.Fatal(http.ListenAndServe(":7788", r))
}

func Grpc() {
	c := config.LoadConfig()
	ls, err := net.Listen(c.User.Host, c.User.Port)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	server := connections.NewService()

	notification.RegisterNotificationServer(s, server)
	reflection.Register(s)

	fmt.Printf("server started on the port %s", c.User.Port)

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
