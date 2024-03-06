package main

import (
	"fmt"
	"net/http"

	// "github.com/ArdynRF/Portofolio-TI/handler"
	"github.com/ArdynRF/Portofolio-TI/routes"
	// "github.com/gin-gonic/gin"
)

func main() {

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
