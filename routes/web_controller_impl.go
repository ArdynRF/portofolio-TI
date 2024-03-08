package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Dashboard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	// ParseFiles akan membaca file template Anda dan mengembalikan *template.Template
	tmpl, err := template.ParseFiles("views/header.html", "views/footer.html", "views/table_perusahaan.html")
	if err != nil {
		// Jangan lupa untuk menangani error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Anda bisa meneruskan data ke Execute untuk digunakan dalam template Anda
	data := map[string]interface{}{
		// Isi dengan data Anda
	}

	err = tmpl.ExecuteTemplate(w, "table_perusahaan.html", data)
	if err != nil {
		// Jangan lupa untuk menangani error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func ListApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// ParseFiles akan membaca file template Anda dan mengembalikan *template.Template
	tmpl, err := template.ParseFiles("views/header.html", "views/footer.html", "views/list.html")
	if err != nil {
		// Jangan lupa untuk menangani error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Anda bisa meneruskan data ke Execute untuk digunakan dalam template Anda
	data := map[string]interface{}{
		// Isi dengan data Anda
	}

	err = tmpl.ExecuteTemplate(w, "list.html", data)
	if err != nil {
		// Jangan lupa untuk menangani error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func TambahApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// ParseFiles akan membaca file template Anda dan mengembalikan *template.Template
	tmpl, err := template.ParseFiles("views/header.html", "views/footer.html", "views/tambahapp.html")
	if err != nil {
		// Jangan lupa untuk menangani error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Anda bisa meneruskan data ke Execute untuk digunakan dalam template Anda
	data := map[string]interface{}{
		// Isi dengan data Anda
	}

	err = tmpl.ExecuteTemplate(w, "tambahapp.html", data)
	if err != nil {
		// Jangan lupa untuk menangani error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
