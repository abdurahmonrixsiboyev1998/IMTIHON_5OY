package handler

import (
	broadcas "api_geteway/internal/broadcast"
	"api_geteway/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	B *broadcas.Adjust
}

// Register godoc
// @Summary Register a new user
// @Description Registers a new user and sends a verification code to their email
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.RegisterUserRequest true "Registration details"
// @Success 200 {string} string "Verification code is sent to your email"
// @Failure 500 {string} string "Internal server error"
// @Router /users/register [post]
func (u *Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.Register(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Verification code is sent to your email")
}

// Verify godoc
// @Summary Verify user account
// @Description Verifies the user account using the provided verification details
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.VerifyRequest true "Verification details"
// @Success 200 {string} string "You have verified your account and now you can log in"
// @Failure 500 {string} string "Internal server error"
// @Router /users/verify [post]
func (u *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.VerifyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := u.B.Verify(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("You have verified your account and now you can log in")
}

// LogIn godoc
// @Summary Log in a user
// @Description Authenticates the user and returns a token if successful
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.LogInRequest true "Login credentials"
// @Success 200 {string} string "JWT token"
// @Failure 500 {string} string "Internal server error"
// @Router /users/login [post]
func (u *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.LogInRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := u.B.Login(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(token)
}

// GetUser godoc
// @Summary Retrieve user information
// @Description Fetches details of a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.GetUserResponse "User details"
// @Failure 500 {string} string "Internal server error"
// @Router /users/ [get]
func (u *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetUser(&models.GetUserRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateUser godoc
// @Summary Update user information
// @Description Updates details of a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body models.UpdateUserRequest true "Updated user details"
// @Success 200 {string} string "Update notification message"
// @Failure 500 {string} string "Internal server error"
// @Router /users/update [put]
func (u *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req models.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.ID = int32(id)
	if err := u.B.UpdateUser(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Your account is updating we'll notify you when it's updated")
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "User deletion message"
// @Failure 500 {string} string "Internal server error"
// @Router /users/delete [delete]
func (u *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.DeleteUser(&models.GetUserRequest{ID: int32(id)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Your account is deleting")
}

// LogOut godoc
// @Summary Log out a user
// @Description Logs out a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "Logout confirmation message"
// @Failure 500 {string} string "Internal server error"
// @Router /users/logout [post]
func (u *Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.Logout(&models.GetUserRequest{ID: int32(id)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("You have logged out!")
}

// Hotel---Service

// CreateHotel godoc
// @Summary Create a new hotel
// @Description Adds a new hotel to the system
// @Tags Hotel
// @Accept json
// @Produce json
// @Param request body models.CreateHotelRequest true "Hotel details"
// @Success 200 {string} string "Hotel created successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels/create [post]
func (u *Handler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateHotelRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.CreateHotel(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Hotel created successfully")
}

// GetHotel godoc
// @Summary Retrieve hotel information
// @Description Fetches details of a hotel by its ID
// @Tags Hotel
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Success 200 {object} models.GetHotelResponse "Hotel details"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels/ [get]
func (u *Handler) GetHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetHotel(&models.GetHotelRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// GetHotels godoc
// @Summary Retrieve a list of hotels
// @Description Fetches a list of all hotels in the system
// @Tags Hotel
// @Accept json
// @Produce json
// @Success 200 {array} models.GetHotelResponse "List of hotels"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels [get]
func (u *Handler) GetHotels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := u.B.GetHotels(&models.GetsRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateHotel godoc
// @Summary Update hotel information
// @Description Updates the details of an existing hotel by its ID
// @Tags Hotel
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Param request body models.UpdateHotelRequest true "Updated hotel details"
// @Success 200 {string} string "Hotel details updated successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels/update [put]
func (u *Handler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req models.UpdateHotelRequest
	req.ID = int32(id)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.UpdateHotel(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Hotel deails are updated")
}

// DeleteHotel godoc
// @Summary Delete a hotel
// @Description Deletes a hotel by its ID
// @Tags Hotel
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Success 200 {string} string "Hotel deletion confirmation"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels/delete [delete]
func (u *Handler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.DeleteHotel(&models.GetHotelRequest{ID: int32(id)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Hotel Deleted")
}

// CreateRoom godoc
// @Summary Create a new room
// @Description Adds a new room to a hotel
// @Tags Room
// @Accept json
// @Produce json
// @Param request body models.CreateRoomRequest true "Room details"
// @Success 200 {string} string "Room created successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels/rooms/create [post]
func (u *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateRoomRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.CreateRoom(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("room created in hotel")
}

// GetRoom godoc
// @Summary Retrieve room information
// @Description Fetches details of a specific room within a hotel by its ID and hotel ID
// @Tags Room
// @Accept json
// @Produce json
// @Param room query int true "Room ID"
// @Param hotel query int true "Hotel ID"
// @Success 200 {object} models.GetRoomResponse "Room details"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels/rooms/ [get]
func (u *Handler) GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomid, err := strconv.Atoi(r.URL.Query().Get("room"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hotelid, err := strconv.Atoi(r.URL.Query().Get("hotel"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := u.B.GetRoom(&models.GetRoomRequest{HotelID: int32(hotelid), ID: int32(roomid)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// GetRooms godoc
// @Summary Retrieve rooms by hotel ID
// @Description Fetches a list of all rooms within a specific hotel
// @Tags Room
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Success 200 {array} models.GetRoomResponse "List of rooms"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels/rooms [get]
func (u *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetRooms(&models.GetRoomRequest{HotelID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateRoom godoc
// @Summary Update room information
// @Description Updates the details of an existing room
// @Tags Room
// @Accept json
// @Produce json
// @Param request body models.UpdateRoomRequest true "Updated room details"
// @Success 200 {string} string "Room details updated successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels/rooms/update [put]
func (u *Handler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.UpdateRoom(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Room details are updated")
}

// DeleteRoom godoc
// @Summary Delete a room
// @Description Deletes a specific room from a hotel by its ID and hotel ID
// @Tags Room
// @Accept json
// @Produce json
// @Param room query int true "Room ID"
// @Param hotel query int true "Hotel ID"
// @Success 200 {string} string "Room deleted successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /hotels/rooms/delete [delete]
func (u *Handler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomid, err := strconv.Atoi(r.URL.Query().Get("room"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hotelid, err := strconv.Atoi(r.URL.Query().Get("hotel"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := u.B.DeleteRoom(&models.GetRoomRequest{HotelID: int32(hotelid), ID: int32(roomid)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Room is deleted")
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Creates a new booking for a hotel room based on the provided details
// @Tags Booking
// @Accept json
// @Produce json
// @Param request body models.BookHotelRequest true "Booking details"
// @Success 200 {object} models.GeneralResponse "Booking created successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /bookings [post]
func (u *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.BookHotelRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.CreateBooking(&req)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(res)
}

// GetBooking godoc
// @Summary Retrieve a booking by ID
// @Description Fetches the details of a specific booking using its ID
// @Tags Booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} models.GetUsersBookResponse "Booking details"
// @Failure 500 {string} string "Internal server error"
// @Router /bookings/ [get]
func (u *Handler) GetBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetBooking(&models.GetUsersBookRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateBooking godoc
// @Summary Update a booking
// @Description Updates the details of an existing booking based on the provided ID and updated information
// @Tags Booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Param request body models.BookHotelUpdateRequest true "Updated booking details"
// @Success 200 {object} models.GeneralResponse "Booking updated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /bookings/update [put]
func (u *Handler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req models.BookHotelUpdateRequest
	req.ID = int32(id)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.UpdateBooking(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// DeleteBooking godoc
// @Summary Delete a booking
// @Description Deletes an existing booking based on the provided booking ID
// @Tags Booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {string} string "Booking successfully deleted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /bookings/delete [delete]
func (u *Handler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.DeleteBooking(&models.CancelRoomRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// CreateWaiting godoc
// @Summary Add a new entry to the waiting list
// @Description Creates a new waiting list entry based on the provided information
// @Tags WaitingList
// @Accept json
// @Produce json
// @Param request body models.CreateWaitingList true "Details for the new waiting list entry"
// @Success 201 {object} models.GeneralResponse "Waiting list entry created successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /waitinglists [post]
func (u *Handler) CreateWaiting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateWaitingList

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.CreateWaitinglist(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)

}

// GetWaiting godoc
// @Summary Retrieve a waiting list entry
// @Description Retrieves details of a specific waiting list entry based on the provided ID
// @Tags WaitingList
// @Accept json
// @Produce json
// @Param id path int true "Waiting List Entry ID"
// @Success 200 {object} models.GetWaitinglistResponse "Details of the waiting list entry"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /waitinglists/ [get]
func (u *Handler) GetWaiting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetWaiting(&models.GetWaitinglistRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateWaiting godoc
// @Summary Update a waiting list entry
// @Description Updates details of a specific waiting list entry based on the provided ID and request body
// @Tags WaitingList
// @Accept json
// @Produce json
// @Param id path int true "Waiting List Entry ID"
// @Param request body models.UpdateWaitingListRequest true "Details for updating the waiting list entry"
// @Success 200 {object} models.GeneralResponse "Updated waiting list entry details"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /waitinglists/update [put]
func (u *Handler) UpdateWaiting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var req models.UpdateWaitingListRequest
	req.ID = int32(id)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := u.B.UpdateWaiting(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// DeleteWaiting godoc
// @Summary Delete a waiting list entry
// @Description Deletes a specific waiting list entry based on the provided ID
// @Tags WaitingList
// @Accept json
// @Produce json
// @Param id path int true "Waiting List Entry ID"
// @Success 200 {string} string "Confirmation message"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /waitinglists/delete [delete]
func (u *Handler) DeleteWaiting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	res, err := u.B.DeleteWaiting(&models.DeleteWaitingList{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}
