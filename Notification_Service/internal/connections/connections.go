package connections

import (
	"context"
	"notify_service/api/handler"
	"notify_service/internal/clients/booking"
	"notify_service/internal/clients/hotel"
	"notify_service/internal/services"
	"sync"

	"github.com/gorilla/websocket"
)

func NewWebSocket() *handler.WebSocket {
	h := hotel.Hotel()
	b := booking.Hotel()

	ctx := context.Background()
	return &handler.WebSocket{
		Map:     make(map[string]*websocket.Conn),
		Mutex:   &sync.Mutex{},
		Hotel:   h,
		Booking: b,
		Ctx:     ctx,
	}
}

func NewService() *services.Service {
	a := NewWebSocket()
	return &services.Service{P: a}
}
