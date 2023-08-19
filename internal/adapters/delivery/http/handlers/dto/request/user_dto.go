package request

type UserDTO struct {
	Name        string `json:"name" example:"New user name"`
	CPF         string `json:"cpf" example:"12345678901"`
	Phone       string `json:"phone" example:"(11) 99999-9999"`
	Email       string `json:"email" example:"johndoe@example.com"`
	Password    string `json:"password" example:"123456"`
	DateOfBirth string `json:"date_of_birth" example:"01/01/2000"`
	Admin       bool   `json:"admin" example:"false"`
}
