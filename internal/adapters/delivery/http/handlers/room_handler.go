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

// @Summary Cria uma nova sala
// @Description Cria uma nova sala com os detalhes fornecidos
// @Tags Sala
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param roomDTO body request.RoomDTO true "Detalhes da sala a ser criada"
// @Success 200 {object} response.InfoResponse "Sala criada com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro de requisição inválida ou criação de sala falhou"
// @Router /room [post]
func (instance RoomHandler) CreateRoom(context echo.Context) error {
	var roomDTO request.RoomDTO
	token := context.Request().Header.Get("Authorization")

	err := context.Bind(&roomDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	roomReceived, err := room.NewBuilder().WithCod(roomDTO.Cod).WithNumber(roomDTO.Number).WithVacancies(roomDTO.Vacancies).WithPrice(roomDTO.Price).Build()
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.CreateRoom(*roomReceived, token)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Room created successfully"})
}

// @Summary Lista todas as salas
// @Description Retorna uma lista de todas as salas disponíveis
// @Tags Sala
// @Produce json
// @Success 200 {array} response.Room
// @Failure 400 {object} response.ErrorResponse "Erro ao listar salas"
// @Router /room [get]
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

// @Summary Obtém detalhes de uma sala por ID
// @Description Retorna os detalhes de uma sala com base no ID fornecido
// @Tags Sala
// @Produce json
// @Param id path string true "ID da sala"
// @Success 200 {object} response.Room
// @Failure 400 {object} response.ErrorResponse "Erro ao obter detalhes da sala"
// @Router /room/{id} [get]
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

// @Summary Obtém detalhes de uma sala por Código
// @Description Retorna os detalhes de uma sala com base no código fornecido
// @Tags Sala
// @Produce json
// @Param cod path string true "Código da sala"
// @Success 200 {object} response.Room
// @Failure 400 {object} response.ErrorResponse "Erro ao obter detalhes da sala"
// @Router /room/{cod} [get]
func (instance RoomHandler) GetRoomByCod(context echo.Context) error {
	cod := context.Param("cod")

	roomReceived, err := instance.service.GetRoomByCod(cod)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.NewRoom(*roomReceived))
}

// @Summary Deleta uma sala por ID
// @Description Deleta uma sala com base no ID fornecido
// @Tags Sala
// @Security bearerAuth
// @Param id path string true "ID da sala"
// @Produce json
// @Success 200 {object} response.InfoResponse "Sala deletada com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao deletar sala"
// @Router /room/{id} [delete]
func (instance RoomHandler) DeleteRoomByID(context echo.Context) error {
	id := context.Param("id")
	token := context.Request().Header.Get("Authorization")

	roomID, err := uuid.Parse(id)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.DeleteRoomByID(roomID, token)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Room deleted successfully"})
}

// @Summary Adiciona uma imagem a uma sala por ID
// @Description Adiciona uma imagem a uma sala com base no ID fornecido
// @Tags Sala
// @Security bearerAuth
// @Param imageDTO body request.ImageDTO true "Detalhes da imagem a ser adicionada"
// @Produce json
// @Success 200 {object} response.InfoResponse "Imagem adicionada com sucesso"
// @Failure 400 {object} response.ErrorResponse "Erro ao adicionar imagem à sala"
// @Router /room/image [post]
func (instance RoomHandler) AddImageToRoomById(context echo.Context) error {
	var imageDTO request.ImageDTO
	token := context.Request().Header.Get("Authorization")

	err := context.Bind(&imageDTO)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = instance.service.AddImageToRoomByRoomID(imageDTO.IdUser, imageDTO.ImageUrl, token)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, response.InfoResponse{Message: "Image added successfully"})
}

// @Summary Obtém detalhes de uma sala com imagens por ID
// @Description Retorna os detalhes de uma sala, incluindo suas imagens, com base no ID fornecido
// @Tags Sala
// @Produce json
// @Param id path string true "ID da sala"
// @Success 200 {object} response.Room
// @Failure 400 {object} response.ErrorResponse "Erro ao obter detalhes da sala com imagens"
// @Router /room/{id}/images [get]
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
