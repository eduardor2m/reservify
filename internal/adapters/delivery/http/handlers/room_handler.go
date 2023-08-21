package handlers

import (
	"fmt"
	"net/http"
	"reservify/internal/adapters/delivery/http/handlers/dto/request"
	"reservify/internal/adapters/delivery/http/handlers/dto/response"
	"reservify/internal/app/entity/room"
	"reservify/internal/app/interfaces/primary"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

func (instance RoomHandler) GetRoomByID(context echo.Context) error {
	id := context.Param("id")

	roomID, err := uuid.Parse(id)

	roomReceived, err := instance.service.GetRoomByID(roomID)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.NewRoom(*roomReceived))
}

func (instance RoomHandler) GetRoomByCod(context echo.Context) error {
	cod := context.Param("cod")

	roomReceived, err := instance.service.GetRoomByCod(cod)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.NewRoom(*roomReceived))
}

func (instance RoomHandler) DeleteRoomByID(context echo.Context) error {
	id := context.Param("id")

	roomID, err := uuid.Parse(id)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.DeleteRoomByID(roomID)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Room deleted successfully"})
}

func (instance RoomHandler) AddImageToRoomById(context echo.Context) error {


	var imageDTO request.ImageDTO

	err := context.Bind(&imageDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	fmt.Println(imageDTO)

	err = instance.service.AddImageToRoomById(imageDTO.IdUser, imageDTO.ImageUrl)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Image added successfully"})
}

func (instance RoomHandler) ListAllRoomsWithImages(context echo.Context) error {
	id := context.Param("id")

	idUUID, err := uuid.Parse(id)

	room, err := instance.service.ListAllImagesByRoomID(idUUID)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	var roomsResponse response.Room

	roomsResponse = *response.NewRoom(*room)

	return context.JSON(http.StatusOK, roomsResponse)
}

func NewRoomHandler(service primary.RoomManager) *RoomHandler {
	return &RoomHandler{
		service: service,
	}
}
