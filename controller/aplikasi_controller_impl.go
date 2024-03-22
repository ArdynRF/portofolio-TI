package controller

import (
	"net/http"
	"strconv"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/aplikasi"
	"github.com/ArdynRF/Portofolio-TI/model/web"
	"github.com/ArdynRF/Portofolio-TI/service"
	"github.com/julienschmidt/httprouter"
)

type AplikasiControllerImpl struct {
	AplikasiService service.AplikasiService
}

func NewAplikasiControllerImpl(userService service.AplikasiService) AplikasiController {
	return &AplikasiControllerImpl{
		AplikasiService: userService,
	}
}

func (controller *AplikasiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aplikasiCreateRequest := aplikasi.AplikasiCreateRequest{}
	helper.BodyToRequest(request, &aplikasiCreateRequest)

	response := controller.AplikasiService.Create(request.Context(), aplikasiCreateRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}
	helper.WriteToBody(writer, webResponse)
}

func (controller *AplikasiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aplikasiUpdateRequest := aplikasi.AplikasiUpdateRequest{}
	helper.BodyToRequest(request, &aplikasiUpdateRequest)

	idAppStr := params.ByName("id_app")
	idApp, err := strconv.ParseInt(idAppStr, 10, 64)
	helper.PanicError(err) // Menangani error dengan fungsi PanicError
	aplikasiUpdateRequest.Id_App = idApp

	response := controller.AplikasiService.Update(request.Context(), aplikasiUpdateRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}
	helper.WriteToBody(writer, webResponse)
}

func (controller *AplikasiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	appIdStr := params.ByName("app_id")
	appId, err := strconv.ParseInt(appIdStr, 10, 64)
	helper.PanicError(err) // Menangani error dengan fungsi PanicError

	controller.AplikasiService.Delete(request.Context(), appId)

	webResponse := web.Response{
		Status: "OK",
	}
	helper.WriteToBody(writer, webResponse)
}


func (controller *AplikasiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIdStr := params.ByName("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	helper.PanicError(err) // Menangani error dengan fungsi PanicError

	response := controller.AplikasiService.FindById(request.Context(), userId)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}
	helper.WriteToBody(writer, webResponse)
}
