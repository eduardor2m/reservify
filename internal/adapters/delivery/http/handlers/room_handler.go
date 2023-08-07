package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reservify/internal/adapters/delivery/http/handlers/dto/request"
	"reservify/internal/adapters/delivery/http/handlers/dto/response"
	"reservify/internal/app/entity/room"
	"reservify/internal/app/interfaces/primary"
)

type RoomHandler struct {
	service primary.RoomManager
}

func (instance RoomHandler) CreateRoom(context echo.Context) error {
	var roomDTO request.RoomDTO

	err := context.Bind(&roomDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	roomReceived, err := room.NewBuilder().WithCod(roomDTO.Cod).WithNumber(roomDTO.Number).WithVacancies(roomDTO.Vacancies).WithPrice(roomDTO.Price).Build()
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.CreateRoom(*roomReceived)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Room created successfully"})
}

func (instance RoomHandler) ListAllRooms(context echo.Context) error {
	rooms, err := instance.service.ListAllRooms()
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var roomsResponse []response.Room

	for _, roomDB := range rooms {
		roomsResponse = append(roomsResponse, *response.NewRoom(roomDB))
	}

	return context.JSON(http.StatusOK, roomsResponse)
}

func NewRoomHandler(service primary.RoomManager) *RoomHandler {
	return &RoomHandler{
		service: service,
	}
}
