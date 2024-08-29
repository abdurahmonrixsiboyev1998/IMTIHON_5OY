package method

import (
	"context"
	"hotel_service/internal/services"
	"hotel_service/internal/protos/hotel"
	"log"
)

type GrpcService struct {
	hotel.UnimplementedHotelServer
	A *services.Adjust
}

func (u *GrpcService) CreateHotel(ctx context.Context, req *hotel.CreateHotelRequest) (*hotel.GeneralResponse, error) {
	r, err := u.A.CreateHotel(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}

func (u *GrpcService) CreateRoom(ctx context.Context, req *hotel.CreateRoomRequest) (*hotel.GeneralResponse, error) {
	r, err := u.A.CreateRoom(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
func (u *GrpcService) DeleteRoom(ctx context.Context, req *hotel.GetroomRequest) (*hotel.GeneralResponse, error) {
	r, err := u.A.DeleteRoom(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
func (u *GrpcService) Delte(ctx context.Context, req *hotel.GetHotelRequest) (*hotel.GeneralResponse, error) {
	r, err := u.A.DeleteHotel(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
func (u *GrpcService) Get(ctx context.Context, req *hotel.GetroomRequest) (*hotel.UpdateRoomRequest, error) {
	r, err := u.A.GetRoom(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
func (u *GrpcService) GetHotel(ctx context.Context, req *hotel.GetHotelRequest) (*hotel.GetHotelResponse, error) {
	r, err := u.A.GetHotel(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
func (u *GrpcService) GetRooms(ctx context.Context, req *hotel.GetroomRequest) (*hotel.GetroomResponse, error) {
	r, err := u.A.GetRooms(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
func (u *GrpcService) Gets(ctx context.Context, req *hotel.GetsRequest) (*hotel.GetsResponse, error) {
	r, err := u.A.Gets(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
func (u *GrpcService) Update(ctx context.Context, req *hotel.UpdateHotelRequest) (*hotel.GeneralResponse, error) {
	r, err := u.A.UpdateHotel(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
func (u *GrpcService) UpdateRoom(ctx context.Context, req *hotel.UpdateRoomRequest) (*hotel.GeneralResponse, error) {
	r, err := u.A.UpdateRooms(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
