package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type TypesRepositoryImpl struct{}

func NewTypesRepositoryImpl() TypesRepository {
	return &TypesRepositoryImpl{}
}

func (repository *TypesRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, tipe entity.Types) entity.Types {
	SQL := "INSERT INTO types (id_status, nama_status) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, tipe.Id_Type, tipe.Nama_Type)
	helper.PanicError(err)
	fmt.Println(tipe)
	return tipe
}
