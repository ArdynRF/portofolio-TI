package web

type UsersUpdateRequest struct {
	Id       string `validate:"required" json:"id"`
	Username string `validate:"required,min=3,max=50" json:"username"`
	Nama     string `validate:"required,min=3,max=50" json:"nama"`
	Role     int64  `validate:"required" json:"role"`
}
