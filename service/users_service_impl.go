package service

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/ArdynRF/Portofolio-TI/exception"
	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
	"github.com/ArdynRF/Portofolio-TI/model/web"
	"github.com/ArdynRF/Portofolio-TI/repository"
	"github.com/ArdynRF/Portofolio-TI/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UsersCreateRequest) web.UsersResponse {
	err := service.validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.Defer(tx)

	passwordHash, err := utils.HashPassword(request.Password)
	helper.PanicError(err)
	user := entity.Users{
		Id:        utils.Uuid(),
		Username:  request.Username,
		Email:     request.Email,
		Password:  passwordHash,
		Nama:      request.Nama,
		Role:      request.Role,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	user = service.UserRepository.Create(ctx, tx, user)
	return utils.UserResponse(user)

}
func (service *UserServiceImpl) Update(ctx context.Context, request web.UsersUpdateRequest) web.UsersResponse {
	err := service.validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.Defer(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user.Username = request.Username
	user.Nama = request.Nama
	user.UpdatedAt = time.Now().Unix()

	user = service.UserRepository.Update(ctx, tx, user)
	return utils.UserResponse(user)

}
func (service *UserServiceImpl) Delete(ctx context.Context, userId string) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.Defer(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicError(err)

	service.UserRepository.Delete(ctx, tx, user)

}
func (service *UserServiceImpl) FindById(ctx context.Context, userId string) web.UsersResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.Defer(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	helper.PanicError(err)
	return utils.UserResponse(user)
}
func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UsersResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.Defer(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	var userResponses []web.UsersResponse
	for _, user := range users {
		userResponses = append(userResponses, utils.UserResponse(user))
	}
	return userResponses
}

func (service *UserServiceImpl) Auth(ctx context.Context, request web.UserAuthRequest) web.TokenResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.Defer(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	jwtExpiredTimeToken, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_TOKEN"))
	jwtExpiredTimeRefreshToken, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_REFRESH_TOKEN"))

	tokenCreateRequest := web.TokenCreateRequest{
		UserId:   user.Id,
		Username: user.Username,
		Email:    user.Email,
		Nama:     user.Nama,
	}
	token := web.TokenResponse{
		Token:        utils.CreateToken(tokenCreateRequest, time.Duration(jwtExpiredTimeToken)),
		RefreshToken: utils.CreateToken(tokenCreateRequest, time.Duration(jwtExpiredTimeRefreshToken)),
	}
	return token
}

func (service *UserServiceImpl) CreateWithRefreshToken(ctx context.Context, refreshToken string) web.TokenResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.Defer(tx)

	claims := utils.ClaimsToken(refreshToken)
	_, err = service.UserRepository.FindById(ctx, tx, claims.UserId)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}
	tokenCreateRequest := web.TokenCreateRequest{
		UserId:   claims.UserId,
		Username: claims.Username,
		Email:    claims.Email,
		Nama:     claims.Nama,
	}
	jwtExpiredTimeToken, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_TOKEN"))
	jwtExpiredTimeRefreshToken, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_REFRESH_TOKEN"))

	token := web.TokenResponse{
		Token:        utils.CreateToken(tokenCreateRequest, time.Duration(jwtExpiredTimeToken)),
		RefreshToken: utils.CreateToken(tokenCreateRequest, time.Duration(jwtExpiredTimeRefreshToken)),
	}
	return token
}
