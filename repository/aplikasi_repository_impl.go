package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type AplikasiRepositoryImpl struct{}

func NewAplikasiRepositoryImpl() AplikasiRepository {
	return &AplikasiRepositoryImpl{}
}
func (repository *AplikasiRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, aplikasi entity.Aplikasi) entity.Aplikasi {
	SQL := "INSERT INTO aplikasi (nama_aplikasi,tahun,alias,url,bussinesowner,acquisition,foto,provider,vendor,fe_language,fe_framework,be_language,be_framework,id_perusahaan,id_category,id_status,id_valuechain,id_asset,id_contract,id_method,id_type,id_architecture,id_infrastructure, created_at,updated_at, review, approval) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, aplikasi.NamaApp, aplikasi.Tahun, aplikasi.Alias, aplikasi.URL, aplikasi.BussinesOwner, aplikasi.FrontEndLang, aplikasi.FrontEndFrame, aplikasi.BackEndLang, aplikasi.BackEndFrame, aplikasi.Acquisition, aplikasi.Foto, aplikasi.Provider, aplikasi.Vendor, aplikasi.Id_Perusahaan, aplikasi.Id_Category, aplikasi.Id_Status, aplikasi.Id_ValueChain, aplikasi.Id_Assets, aplikasi.Id_Contract, aplikasi.Id_Method, aplikasi.Id_Type, aplikasi.Id_Architecture, aplikasi.Id_Infrastructure, aplikasi.Created_At, aplikasi.Updated_At, aplikasi.Review, aplikasi.Approval)
	helper.PanicError(err)
	fmt.Println(aplikasi)
	return aplikasi
}
func (repository *AplikasiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, aplikasi entity.Aplikasi) {
	SQL := "DELETE FROM aplikasi WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, aplikasi.Id_App)
	helper.PanicError(err)

}
func (repository *AplikasiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, appId int64) (entity.Aplikasi, error) {
	SQL := "SELECT id_aplikasi, nama_aplikasi,tahun,alias,url,bussinesowner,acquisition,foto,provider,vendor,fe_language,fe_framework,be_language,be_framework,id_perusahaan,id_category,id_status,id_valuechain,id_asset,id_contract,id_method,id_type,id_architecture,id_infrastructure, created_at,updated_at, review, approval FROM aplikasi WHERE id_aplikasi=?"
	rows, err := tx.QueryContext(ctx, SQL, appId)
	helper.PanicError(err)
	defer rows.Close()
	aplikasi := entity.Aplikasi{}
	if rows.Next() {
		err := rows.Scan(&aplikasi.NamaApp, &aplikasi.Tahun, &aplikasi.URL, &aplikasi.BussinesOwner, &aplikasi.Acquisition, aplikasi.Foto, &aplikasi.Provider, &aplikasi.Vendor, &aplikasi.FrontEndLang, &aplikasi.FrontEndFrame, &aplikasi.BackEndLang, &aplikasi.BackEndFrame)
		helper.PanicError(err)
		return aplikasi, nil
	} else {
		return aplikasi, errors.New("user not found")
	}
}


func (repository *AplikasiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, aplikasi entity.Aplikasi) entity.Aplikasi {
	SQL := `UPDATE aplikasi SET nama_aplikasi = ?, tahun = ?, alias = ?, url = ?, bussinesowner = ?, acquisition = ?, foto = ?, provider = ?, vendor = ?, fe_language = ?, fe_framework = ?, be_language = ?, be_framework = ?, id_perusahaan = ?, id_category = ?, id_status = ?, id_valuechain = ?, id_asset = ?, id_contract = ?, id_method = ?, id_type = ?, id_architecture = ?, id_infrastructure = ?, updated_at = ?, review = ?, approval = ? WHERE id_aplikasi = ?`
	_, err := tx.ExecContext(ctx, SQL, aplikasi.NamaApp, aplikasi.Tahun, aplikasi.Alias, aplikasi.URL, aplikasi.BussinesOwner, aplikasi.FrontEndLang, aplikasi.FrontEndFrame, aplikasi.BackEndLang, aplikasi.BackEndFrame, aplikasi.Acquisition, aplikasi.Foto, aplikasi.Provider, aplikasi.Vendor, aplikasi.Id_Perusahaan, aplikasi.Id_Category, aplikasi.Id_Status, aplikasi.Id_ValueChain, aplikasi.Id_Assets, aplikasi.Id_Contract, aplikasi.Id_Method, aplikasi.Id_Type, aplikasi.Id_Architecture, aplikasi.Id_Infrastructure, aplikasi.Updated_At, aplikasi.Review, aplikasi.Approval, aplikasi.Id_App)
	helper.PanicError(err)
	// fmt.Println(aplikasi)
	return aplikasi
}