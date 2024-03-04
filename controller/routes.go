package controller

import (
	"net/http"
	"text/template"
)

func Dashboard(response http.ResponseWriter, request *http.Request) {

	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)

}
func Perusahaan(response http.ResponseWriter, request *http.Request) {

	temp, err := template.ParseFiles("views/table_perusahaan.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)

}

func ListApp(response http.ResponseWriter, request *http.Request) {

	temp, err := template.ParseFiles("views/list.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)

}

func TambahApp(response http.ResponseWriter, request *http.Request) {

	temp, err := template.ParseFiles("views/tambahapp.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)

}
