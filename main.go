package main

import (
	"fmt"
	"net/http"

	"github.com/ArdynRF/Portofolio-TI/config"
	"github.com/ArdynRF/Portofolio-TI/controller"
	"github.com/ArdynRF/Portofolio-TI/exception"
	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/middleware"
	"github.com/ArdynRF/Portofolio-TI/repository"
	"github.com/ArdynRF/Portofolio-TI/routes"
	"github.com/ArdynRF/Portofolio-TI/service"
	"github.com/joho/godotenv"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//load environment
	envErr := godotenv.Load(".env")
	helper.PanicError(envErr)

	//db
	db := config.Database()

	//validation
	validate := validator.New()

	//repo
	userRepository := repository.NewUserRepositoryImpl()
	aplikasiRepository := repository.NewAplikasiRepositoryImpl()

	//service
	userService := service.NewUserServiceImpl(userRepository, db, validate)
	aplikasiService := service.NewAplikasiServiceImpl(aplikasiRepository, db, validate)

	//user controller
	userController := controller.NewUserControllerImpl(userService)
	aplikasiController := controller.NewAplikasiControllerImpl(aplikasiService)

	// initialize router
	router := httprouter.New()

	//router
	// [USER]
	router.POST("/api/v1/user", userController.Create)
	router.POST("/api/v1/auth", userController.Auth)
	router.POST("/api/v1/refresh-token", userController.CreateWithRefreshToken)
	router.PUT("/api/v1/user/:user_id", userController.Update)
	router.DELETE("/api/v1/user/:user_id", userController.Delete)
	router.GET("/api/v1/user/:user_id", userController.FindById)
	router.GET("/api/v1/user", userController.FindAll)
	router.POST("/api/v1/aplikasi", aplikasiController.Create)

	// [WEB]
	router.GET("/", routes.Login)
	router.GET("/table_perusahaan", routes.Perusahaan)
	router.GET("/listapp", routes.ListApp)
	router.GET("/tambahapp", routes.TambahApp)
	router.GET("/dashboard", routes.Dashboard)

	router.PanicHandler = exception.ErrorHandler

	router.ServeFiles("/assets/*filepath", http.Dir("assets")) // ASsets Static

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicError(err)

	fmt.Println("Hello World")

}
