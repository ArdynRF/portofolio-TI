package web

type UsersResponse struct {
	Id        string `json:"id"`
	Username  string ` json:"username"`
	Email     string ` json:"email"`
	Password  string ` json:"password"`
	Nama      string ` json:"nama"`
	Role      int64  `json:role`
	CreatedAt int64  `json:created_at"`
	UpdatedAt int64  `json:updated_at`
	// createdAt time.Time `json:created'

}
