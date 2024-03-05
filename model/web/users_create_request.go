package web

type UsersCreateRequest struct {
	Username string `validate:"required,min=3,max=50" json:"username"`
	Email    string `validate:"required,min=3,max=50" json:"email"`
	Password string `validate:"required" json:"password"`
	Nama     string `validate:"required,min=3,max=50" json:"nama"`
	Role     int64  `validate:"required", json:role`
}
