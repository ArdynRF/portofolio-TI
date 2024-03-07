package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Dashboard(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// ParseFiles akan membaca file template Anda dan mengembalikan *template.Template
	tmpl, err := template.ParseFiles("views/header.html", "views/footer.html", "views/index.html")
	if err != nil {
		// Jangan lupa untuk menangani error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Anda bisa meneruskan data ke Execute untuk digunakan dalam template Anda
	data := map[string]interface{}{
		// Isi dengan data Anda
	}

	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		// Jangan lupa untuk menangani error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
