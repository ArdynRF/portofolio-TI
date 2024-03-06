package utils

import (
	"github.com/ArdynRF/Portofolio-TI/model/entity"
	"github.com/ArdynRF/Portofolio-TI/model/web"
)

func UserResponse(user entity.Users) web.UsersResponse {
	return web.UsersResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Nama:      user.Nama,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
