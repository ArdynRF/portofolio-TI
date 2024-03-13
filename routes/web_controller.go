package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type WebController interface {
	Dashboard(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Perusahaan(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	ListApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	TambahApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}
