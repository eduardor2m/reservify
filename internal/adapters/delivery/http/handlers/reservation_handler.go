package handlers

import (
	"net/http"
	"reservify/internal/adapters/delivery/http/handlers/dto/request"
	"reservify/internal/adapters/delivery/http/handlers/dto/response"
	"reservify/internal/app/entity/reservation"
	"reservify/internal/app/interfaces/primary"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	service primary.ReservationManager
}

func (instance ReservationHandler) CreateReservation(context echo.Context) error {
	var reservationDTO request.ReservationDTO
	token := context.Request().Header.Get("Authorization")

	err := context.Bind(&reservationDTO)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	reservationReceived, err := reservation.NewBuilder().WithIdRoom(reservationDTO.IdRoom).WithIdUser(reservationDTO.IdUser).WithCheckIn(reservationDTO.CheckIn).WithCheckOut(reservationDTO.CheckOut).Build()

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.CreateReservation(*reservationReceived, token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Room rented successfully"})
}

func (instance ReservationHandler) ListAllReservations(context echo.Context) error {
	token := context.Request().Header.Get("Authorization")

	reservations, err := instance.service.ListAllReservations(token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var reservationsResponse []response.Reservation

	for _, reservationDB := range reservations {
		reservationsResponse = append(reservationsResponse, *response.NewReservation(reservationDB))
	}

	if len(reservationsResponse) == 0 {
		return context.JSON(http.StatusOK, []response.Reservation{})
	}

	return context.JSON(http.StatusOK, reservationsResponse)
}

func (instance ReservationHandler) GetReservationByID(context echo.Context) error {
	id := context.Param("id")
	token := context.Request().Header.Get("Authorization")

	reservationID, err := uuid.Parse(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	reservationReceived, err := instance.service.GetReservationByID(reservationID, token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.NewReservation(*reservationReceived))
}

func (instance ReservationHandler) GetReservationsByRoomID(context echo.Context) error {
	id := context.Param("id_room")
	token := context.Request().Header.Get("Authorization")

	reservationID, err := uuid.Parse(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	reservationsReceived, err := instance.service.GetReservationsByRoomID(reservationID, token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var reservationsResponse []response.Reservation

	for _, reservationDB := range reservationsReceived {
		reservationsResponse = append(reservationsResponse, *response.NewReservation(reservationDB))
	}

	if len(reservationsResponse) == 0 {
		return context.JSON(http.StatusOK, []response.Reservation{})
	}

	return context.JSON(http.StatusOK, reservationsResponse)
}

func (instance ReservationHandler) GetReservationsByUserID(context echo.Context) error {
	id := context.Param("id_user")

	reservationID, err := uuid.Parse(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	reservationsReceived, err := instance.service.GetReservationsByUserID(reservationID)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var reservationsResponse []response.Reservation

	for _, reservationDB := range reservationsReceived {
		reservationsResponse = append(reservationsResponse, *response.NewReservation(reservationDB))
	}

	if len(reservationsResponse) == 0 {
		return context.JSON(http.StatusOK, []response.Reservation{})
	}

	return context.JSON(http.StatusOK, reservationsResponse)

}

func (instance ReservationHandler) DeleteReservationByID(context echo.Context) error {
	id := context.Param("id")
	token := context.Request().Header.Get("Authorization")

	reservationID, err := uuid.Parse(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.DeleteReservationByID(reservationID, token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Reservation deleted successfully"})

}

func NewReservationHandler(service primary.ReservationManager) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}
