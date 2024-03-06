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

	//service
	userService := service.NewUserServiceImpl(userRepository, db, validate)

	//user controller
	userController := controller.NewUserControllerImpl(userService)

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

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicError(err)

	fmt.Println("Hello World")

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// router := gin.Default()

	// router.GET("/", handler.RootHandler)

	// router.Run(":3000")

	http.HandleFunc("/", routes.Dashboard)
	http.HandleFunc("/table_perusahaan", routes.Perusahaan)
	http.HandleFunc("/listapp", routes.ListApp)
	http.HandleFunc("/tambahapp", routes.TambahApp)

	http.ListenAndServe(":3000", nil)

}
