package handlers

import (
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

	if len(roomsResponse) == 0 {
		return context.JSON(http.StatusOK, []response.Room{})
	}

	return context.JSON(http.StatusOK, roomsResponse)
}

func (instance RoomHandler) GetRoomByID(context echo.Context) error {
	id := context.Param("id")

	roomID, err := uuid.Parse(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

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

	err = instance.service.AddImageToRoomByRoomID(imageDTO.IdUser, imageDTO.ImageUrl)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Image added successfully"})
}

func (instance RoomHandler) GetRoomWithImages(context echo.Context) error {
	id := context.Param("id")

	idUUID, err := uuid.Parse(id)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	room, err := instance.service.GetRoomWithImagesByRoomID(idUUID)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	roomsResponse := *response.NewRoom(*room)

	if roomsResponse.Images == nil {
		return context.JSON(http.StatusOK, response.InfoResponse{Message: "Room has no images"})
	}

	return context.JSON(http.StatusOK, roomsResponse)
}

func NewRoomHandler(service primary.RoomManager) *RoomHandler {
	return &RoomHandler{
		service: service,
	}
}
