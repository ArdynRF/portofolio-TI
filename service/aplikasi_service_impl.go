package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/aplikasi"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
	"github.com/ArdynRF/Portofolio-TI/repository"
	"github.com/ArdynRF/Portofolio-TI/utils"
	"github.com/go-playground/validator/v10"
)

type AplikasiServiceImpl struct {
	AplikasiRepository repository.AplikasiRepository
	DB                 *sql.DB
	validate           *validator.Validate
}

func NewAplikasiServiceImpl(aplikasiRepository repository.AplikasiRepository, DB *sql.DB, validate *validator.Validate) AplikasiService {
	return &AplikasiServiceImpl{
		AplikasiRepository: aplikasiRepository,
		DB:                 DB,
		validate:           validate,
	}
}

func (service *AplikasiServiceImpl) Create(ctx context.Context, request aplikasi.AplikasiCreateRequest) aplikasi.AplikasiResponse {
	err := service.validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.Defer(tx)

	helper.PanicError(err)
	aplikasi := entity.Aplikasi{
		NamaApp:           request.NamaApp,
		Tahun:             request.Tahun,
		Alias:             request.Alias,
		URL:               request.URL,
		BussinesOwner:     request.BussinesOwner,
		FrontEndLang:      request.FrontEndLang,
		FrontEndFrame:     request.FrontEndFrame,
		BackEndLang:       request.BackEndLang,
		BackEndFrame:      request.BackEndFrame,
		Acquisition:       request.Acquisition,
		Foto:              request.Foto,
		Provider:          request.Provider,
		Vendor:            request.Vendor,
		Id_Perusahaan:     request.Id_Perusahaan,
		Id_Category:       request.Id_Category,
		Id_Status:         request.Id_Status,
		Id_ValueChain:     request.Id_ValueChain,
		Id_Assets:         request.Id_Assets,
		Id_Contract:       request.Id_Contract,
		Id_Method:         request.Id_Method,
		Id_Type:           request.Id_Type,
		Id_Architecture:   request.Id_Architecture,
		Id_Infrastructure: request.Id_Infrastructure,
		Created_At:        time.Now().Unix(),
		Updated_At:        time.Now().Unix(),
	}
	aplikasi = service.AplikasiRepository.Create(ctx, tx, aplikasi)
	return utils.AplikasiResponse(aplikasi)

}
