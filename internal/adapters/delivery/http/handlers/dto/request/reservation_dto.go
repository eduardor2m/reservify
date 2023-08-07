package request

type ReservationDTO struct {
	IdUser   string `json:"id_user"`
	IdRoom   string `json:"id_room"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
}
