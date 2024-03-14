package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type AplikasiRepositoryImpl struct{}

func NewAplikasiRepositoryImpl() AplikasiRepository {
	return &AplikasiRepositoryImpl{}
}
func (repository *AplikasiRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, aplikasi entity.Aplikasi) entity.Aplikasi {
	SQL := "INSERT INTO aplikasi (nama_aplikasi,tahun,alias,url,bussinesowner,acquisition,foto,provider,vendor,fe_language,fe_framework,be_language,be_framework,id_perusahaan,id_category,id_status,id_valuechain,id_asset,id_contract,id_method,id_type,id_architecture,id_infrastructure, created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, aplikasi.NamaApp, aplikasi.Tahun, aplikasi.Alias, aplikasi.URL, aplikasi.BussinesOwner, aplikasi.FrontEndLang, aplikasi.FrontEndFrame, aplikasi.BackEndLang, aplikasi.BackEndFrame, aplikasi.Acquisition, aplikasi.Foto, aplikasi.Provider, aplikasi.Vendor, aplikasi.Id_Perusahaan, aplikasi.Id_Category, aplikasi.Id_Status, aplikasi.Id_ValueChain, aplikasi.Id_Assets, aplikasi.Id_Contract, aplikasi.Id_Method, aplikasi.Id_Type, aplikasi.Id_Architecture, aplikasi.Id_Infrastructure, aplikasi.Created_At, aplikasi.Updated_At)
	helper.PanicError(err)
	fmt.Println(aplikasi)
	return aplikasi
}
