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

// @Summary Cria uma nova reserva de sala
// @Description Cria uma nova reserva de sala com base nos detalhes fornecidos
// @Tags Reserva
// @Produce json
// @Param Authorization header string true "Token de autenticação do usuário"
// @Param reservationDTO body request.ReservationDTO true "Detalhes da reserva a ser criada"
// @Success 200 {object} response.InfoResponse "Reserva de sala realizada com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao criar reserva de sala"
// @Router /reservations [post]
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

// @Summary Cria uma nova reserva de sala para o usuário autenticado
// @Description Cria uma nova reserva de sala com base nos detalhes fornecidos pelo usuário autenticado
// @Tags Reserva
// @Produce json
// @Param Authorization header string true "Token de autenticação do usuário"
// @Param reservationDTO body request.ReservationDTO true "Detalhes da reserva a ser criada"
// @Success 200 {object} response.InfoResponse "Reserva de sala realizada com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao criar reserva de sala"
// @Router /reservations/my [post]
func (instance ReservationHandler) CreateMyReservation(context echo.Context) error {
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

	err = instance.service.CreateMyReservation(*reservationReceived, token)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Room rented successfully"})
}

// @Summary Lista todas as reservas de salas
// @Description Retorna uma lista de todas as reservas de salas
// @Tags Reserva
// @Produce json
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {array} response.Reservation
// @Failure 400 {object} response.ErrorResponse "Erro ao listar reservas de salas"
// @Router /reservations [get]
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

// @Summary Obtém detalhes de uma reserva de sala por ID
// @Description Retorna os detalhes de uma reserva de sala com base no ID fornecido
// @Tags Reserva
// @Produce json
// @Param id path string true "ID da reserva"
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {object} response.Reservation
// @Failure 400 {object} response.ErrorResponse "Erro ao obter detalhes da reserva de sala"
// @Router /reservations/{id} [get]
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

// @Summary Obtém as reservas de salas por ID de sala
// @Description Retorna as reservas de salas com base no ID da sala fornecido
// @Tags Reserva
// @Produce json
// @Param id_room path string true "ID da sala"
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {array} response.Reservation
// @Failure 400 {object} response.ErrorResponse "Erro ao obter reservas de salas por ID de sala"
// @Router /reservations/room/{id_room} [get]
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

// @Summary Obtém as reservas de salas por ID de usuário
// @Description Retorna as reservas de salas com base no ID do usuário fornecido
// @Tags Reserva
// @Produce json
// @Param id_user path string true "ID do usuário"
// @Success 200 {array} response.Reservation
// @Failure 400 {object} response.ErrorResponse "Erro ao obter reservas de salas por ID de usuário"
// @Router /reservations/user/{id_user} [get]
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

// @Summary Deleta uma reserva de sala por ID
// @Description Deleta uma reserva de sala com base no ID fornecido
// @Tags Reserva
// @Produce json
// @Param id path string true "ID da reserva"
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {object} response.InfoResponse "Reserva de sala deletada com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao deletar reserva de sala"
// @Router /reservations/{id} [delete]
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

// @Summary Deleta uma reserva de sala do usuário autenticado por ID
// @Description Deleta uma reserva de sala do usuário autenticado com base no ID fornecido
// @Tags Reserva
// @Produce json
// @Param id path string true "ID da reserva"
// @Param Authorization header string true "Token de autenticação do usuário"
// @Success 200 {object} response.InfoResponse "Reserva de sala deletada com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao deletar reserva de sala"
// @Router /reservations/my/{id} [delete]
func (instance ReservationHandler) DeleteMyReservationByID(context echo.Context) error {
	id := context.Param("id")
	token := context.Request().Header.Get("Authorization")

	reservationID, err := uuid.Parse(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.DeleteMyReservationByID(reservationID, token)

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
