package exception

import (
	"net/http"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if badRequest(writer, request, err) {
		return
	} else if notFound(writer, request, err) {
		return
	}
	if unauthorized(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)

}
func badRequest(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		response := web.Response{
			Status: "Bad_Request",
			Data:   exception.Error,
		}
		helper.WriteToBody(writer, response)
		return true
	} else {
		return false
	}
}

func notFound(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		response := web.Response{
			Status: "Not_Found",
			Data:   exception.Error,
		}
		helper.WriteToBody(writer, response)
		return true
	} else {
		return false
	}
}

func unauthorized(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		response := web.Response{
			Status: "Unauthorized",
			Data:   exception.Error,
		}
		helper.WriteToBody(writer, response)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	response := web.Response{
		Status: "Bad_Request",
		Data:   err,
	}
	helper.WriteToBody(writer, response)
}
