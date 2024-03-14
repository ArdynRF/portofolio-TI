package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type PerusahaanRepositoryImpl struct{}

func NewPerusahaanRepositoryImpl() PerusahaanRepository {
	return &PerusahaanRepositoryImpl{}
}
func (repository *PerusahaanRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, perusahaan entity.Perusahaan) entity.Perusahaan {
	SQL := "INSERT INTO perusahaan (id_modul, nama_modul, id_aplikasi) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, perusahaan.Id_Perusahaan, perusahaan.Nama_Perusahaan)
	helper.PanicError(err)
	fmt.Println(perusahaan)
	return perusahaan
}
