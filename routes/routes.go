package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Dashboard(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)

}

func Perusahaan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	temp, err := template.ParseFiles("views/table_perusahaan.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)

}

func ListApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	temp, err := template.ParseFiles("views/list.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)

}

func TambahApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	temp, err := template.ParseFiles("views/tambahapp.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)

}
