package server

import (
	"context"
	"errors"
	"hotel_service/internal/services"
	"hotel_service/internal/models"
	"hotel_service/internal/protos/hotel"
	"log"
)

type Adust struct {
	S *services.Database
}

func (u *Adust) CreateHotel(ctx context.Context, r *hotel.CreateHotelRequest) (*hotel.GeneralResponse, error) {
	if r.Address != "" && r.Location != "" && r.Name != "" && r.Rating != 0 {
		var newreq = models.CreateHotelRequest{
			Name:     r.Name,
			Location: r.Location,
			Rating:   r.Rating,
			Address:  r.Address,
		}
		res, err := u.S.CreateHotel(ctx, &newreq)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &hotel.GeneralResponse{Message: res.Message}, nil
	}
	return nil, errors.New("missing fields")

}
func (u *Adust) GetHotel(ctx context.Context, r *hotel.GetHotelRequest) (*hotel.GetHotelResponse, error) {
	res, err := u.S.GetHotel(ctx, &models.GetHotelRequest{ID: r.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var rooms []*hotel.UpdateRoomRequest
	for _, v := range res.Rooms {
		var all = hotel.UpdateRoomRequest{
			Available:     v.Available,
			HotelId:       v.HotelID,
			RoomType:      v.RoomType,
			Id:            v.ID,
			PricePerNight: v.PricePerNight,
		}
		rooms = append(rooms, &all)
	}
	return &hotel.GetHotelResponse{Id: res.ID, Name: res.Name, Location: res.Location, Rating: res.Rating, Address: res.Address, Rooms: rooms}, nil
}

func (u *Adust) Gets(ctx context.Context, r *hotel.GetsRequest) (*hotel.GetsResponse, error) {
	res, err := u.S.Gets(ctx, &models.GetsRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var hotels []*hotel.UpdateHotelRequest

	for _, v := range res {
		var all = hotel.UpdateHotelRequest{
			Id:       v.ID,
			Name:     v.Name,
			Location: v.Location,
			Rating:   v.Rating,
			Address:  v.Address,
		}
		hotels = append(hotels, &all)
	}
	return &hotel.GetsResponse{Hotels: hotels}, nil
}

func (u *Adust) UpdateHotel(ctx context.Context, r *hotel.UpdateHotelRequest) (*hotel.GeneralResponse, error) {
		var newreq = models.UpdateHotelRequest{
			ID:       r.Id,
			Name:     r.Name,
			Location: r.Location,
			Rating:   r.Rating,
			Address:  r.Address,
		}
		res, err := u.S.UpdateHotel(ctx, &newreq)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &hotel.GeneralResponse{Message: res.Message}, nil
}

func (u *Adust) DeleteHotel(ctx context.Context, r *hotel.GetHotelRequest) (*hotel.GeneralResponse, error) {
	res, err := u.S.DeleteHotel(ctx, &models.GetHotelRequest{ID: r.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &hotel.GeneralResponse{Message: res.Message}, nil
}

func (u *Adust) CreateRoom(ctx context.Context, r *hotel.CreateRoomRequest) (*hotel.GeneralResponse, error) {
	if r.HotelId!=0 && r.RoomType!="" && r.PricePerNight!=0{
		var newrew = models.CreateRoomRequest{
			HotelID:       r.HotelId,
			RoomType:      r.RoomType,
			PricePerNight: r.PricePerNight,
		}
		res, err := u.S.CreateRoom(ctx, &newrew)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &hotel.GeneralResponse{Message: res.Message}, nil
	}
	return nil,errors.New("missing field")
}

func (u *Adust) GetRoom(ctx context.Context, r *hotel.GetroomRequest) (*hotel.UpdateRoomRequest, error) {
	res, err := u.S.GetRoom(ctx, &models.GetRoomRequest{HotelID: r.HotelId, ID: r.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &hotel.UpdateRoomRequest{HotelId: res.HotelID, Id: res.ID, RoomType: res.RoomType, PricePerNight: res.PricePerNight, Available: res.Available}, nil
}

func (u *Adust) GetRooms(ctx context.Context, r *hotel.GetroomRequest) (*hotel.GetroomResponse, error) {
	res, err := u.S.GetRooms(ctx, &models.GetRoomRequest{HotelID: r.HotelId})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var rooms []*hotel.UpdateRoomRequest

	for _, v := range res.Rooms {
		var all = hotel.UpdateRoomRequest{
			HotelId:       v.HotelID,
			Id:            v.ID,
			RoomType:      v.RoomType,
			PricePerNight: v.PricePerNight,
			Available:     v.Available,
		}
		rooms = append(rooms, &all)
	}
	return &hotel.GetroomResponse{Rooms: rooms}, nil
}
func (u *Adust) UpdateRooms(ctx context.Context, r *hotel.UpdateRoomRequest) (*hotel.GeneralResponse, error) {
	var newreq = models.UpdateRoomRequest{
		HotelID:       r.HotelId,
		ID:            r.Id,
		RoomType:      r.RoomType,
		PricePerNight: r.PricePerNight,
		Available:     r.Available,
	}
	res, err := u.S.UpdateRooms(ctx, &newreq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &hotel.GeneralResponse{Message: res.Message}, nil
}

func (u *Adust) DeleteRoom(ctx context.Context, r *hotel.GetroomRequest) (*hotel.GeneralResponse, error) {
	res, err := u.S.DeleteRoom(ctx, &models.GetRoomRequest{HotelID: r.HotelId, ID: r.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &hotel.GeneralResponse{Message: res.Message}, nil
}
