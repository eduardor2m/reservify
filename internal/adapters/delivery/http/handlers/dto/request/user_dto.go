package request

type UserDTO struct {
	Name        string `json:"name" example:"New user name"`
	Email       string `json:"email" example:"johndoe@example.com"`
	Password    string `json:"password" example:"123456"`
	DateOfBirth string `json:"date_of_birth" example:"1990-01-01"`
	Admin       bool   `json:"admin" example:"false"`
}
