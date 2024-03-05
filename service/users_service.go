package service

import (
	"context"

	"github.com/ArdynRF/Portofolio-TI/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UsersCreateRequest) web.UsersResponse
	Update(ctx context.Context, request web.UsersUpdateRequest) web.UsersResponse
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) web.UsersResponse
	// FindbyEmail()
	FindAll(ctx context.Context) []web.UsersResponse
	Auth(ctx context.Context, request web.UserAuthRequest) web.TokenResponse
	CreateWithRefreshToken(ctx context.Context, refreshToken string) web.TokenResponse
}
