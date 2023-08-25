package request

type RoomDTO struct {
	Name        string  `json:"name" example:"New room name"`
	Description string  `json:"description" example:"New room description"`
	Cod         string  `json:"cod" example:"New room cod"`
	Number      int     `json:"number" example:"1"`
	Vacancies   int     `json:"vacancies" example:"2"`
	Price       float64 `json:"price" example:"100.00"`
}
