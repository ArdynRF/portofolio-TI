package web

type TokenCreateRequest struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Nama     string `json:"nama"`
}
