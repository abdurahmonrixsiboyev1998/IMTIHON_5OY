package db_sql

import (
	"fmt"
	"hotel_service/internal/models"
	"log"

	"github.com/Masterminds/squirrel"
)

func GetHotel(req *models.GetHotelRequest) (string, []interface{}, error) {
	q, a, err := squirrel.Select("*").
		From("hotels").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
		
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}

func GetsHotel(req *models.GetsRequest) (string, []interface{}, error) {
	q, a, err := squirrel.Select("*").
		From("hotels").
		ToSql()

	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}

func CreateHotel(req *models.CreateHotelRequest) (string, []interface{}, error) {
	q, a, err := squirrel.Insert("hotels").
		Columns("name", "location", "rating", "address").
		Values(req.Name, req.Location, req.Rating, req.Address).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}


func UpdateHotel(r *models.UpdateHotelRequest) (string, []interface{}, error) {
	setMap := make(map[string]interface{})

	if r.Name != "" {
		setMap["name"] = r.Name
	}

	if r.Location != "" {
		setMap["location"] = r.Location
	}

	if r.Rating != 0 {
		setMap["rating"] = r.Rating
	}

	if r.Address != "" {
		setMap["address"] = r.Address
	}

	if len(setMap) == 0 {
		return "", nil, fmt.Errorf("no fields to update")
	}

	q, a, err := squirrel.Update("hotels").
		SetMap(setMap).
		Where(squirrel.Eq{"id": r.ID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}

func DeleteHotel(req *models.GetHotelRequest) (string, []interface{}, error) {
	q, a, err := squirrel.Delete("hotels").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}

func GetRoom(req *models.GetRoomRequest) (string, []interface{}, error) {
	q, a, err := squirrel.Select("*").
		From("rooms").
		Where(squirrel.Eq{"id": req.ID, "hotel_id": req.HotelID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}

func GetsRoom(req *models.GetRoomRequest) (string, []interface{}, error) {
	q, a, err := squirrel.Select("*").
		From("rooms").
		Where(squirrel.Eq{"available": true, "hotel_id": req.HotelID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}

func CreateRoom(req *models.CreateRoomRequest) (string, []interface{}, error) {
	q, a, err := squirrel.Insert("rooms").
		Columns("hotel_id", "room_type", "price_per_night").
		Values(req.HotelID, req.RoomType, req.PricePerNight).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}

func UpdateRoom(req *models.UpdateRoomRequest) (string, []interface{}, error) {
	setMap := make(map[string]interface{})

	if req.Available{
		setMap["available"] = req.Available
	}else if !req.Available{
		setMap["available"] = req.Available
	}

	if req.RoomType != ""{
		setMap["room_type"] = req.RoomType
	}

	if req.PricePerNight != 0 {
		setMap["price_per_night"] = req.PricePerNight
	}

	if len(setMap) == 0 {
		return "", nil, fmt.Errorf("no fields to update")
	}

	q, a, err := squirrel.Update("rooms").
		SetMap(setMap).
		Where(squirrel.Eq{"id": req.ID, "hotel_id": req.HotelID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}

func DeleteRoom(req *models.GetRoomRequest) (string, []interface{}, error) {
	q, a, err := squirrel.Delete("rooms").
		Where(squirrel.Eq{"id": req.ID, "hotel_id": req.HotelID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}

func GetRoomForHotel(req int) (string, []interface{}, error) {
	q, a, err := squirrel.Select("*").
		From("rooms").
		Where(squirrel.Eq{"hotel_id": req}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	return q, a, nil
}
