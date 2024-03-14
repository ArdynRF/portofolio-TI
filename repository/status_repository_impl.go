package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type StatusRepositoryImpl struct{}

func NewStatusRepositoryImpl() StatusRepository {
	return &StatusRepositoryImpl{}
}

func (repository *StatusRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, status entity.Status) entity.Status {
	SQL := "INSERT INTO status (id_status, nama_status) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, status.Id_Status, status.Nama_Status)
	helper.PanicError(err)
	fmt.Println(status)
	return status
}
