package controller

import (
	"net/http"

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
	userCreateRequest := aplikasi.AplikasiCreateRequest{}
	helper.BodyToRequest(request, &userCreateRequest)

	response := controller.AplikasiService.Create(request.Context(), userCreateRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}
	helper.WriteToBody(writer, webResponse)
}
