package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type ModulsRepositoryImpl struct{}

func NewModulsRepositoryImpl() ModulsRepository {
	return &ModulsRepositoryImpl{}
}
func (repository *ModulsRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, modul entity.Moduls) entity.Moduls {
	SQL := "INSERT INTO moduls (id_modul, nama_modul, id_aplikasi) VALUES (?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, modul.Id_Modul, modul.Nama_Modul, modul.Id_aplikasi)
	helper.PanicError(err)
	fmt.Println(modul)
	return modul
}
