package main

import (
	"fmt"
	"net/http"

	"github.com/ArdynRF/Portofolio-TI/controller"
)

func main() {

	fmt.Println("Hello World")

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	
	http.HandleFunc("/", controller.Dashboard)
	http.HandleFunc("/table_perusahaan", controller.Perusahaan)
	http.HandleFunc("/listapp", controller.ListApp)
	http.HandleFunc("/tambahapp", controller.TambahApp)

	http.ListenAndServe(":3000", nil)

}
