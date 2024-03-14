package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type ValuechainRepositoryImpl struct{}

func NewValuechainRepositoryImpl() ValuechainRepository {
	return &ValuechainRepositoryImpl{}
}

func (repository *ValuechainRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, valuechain entity.Valuechain) entity.Valuechain {
	SQL := "INSERT INTO valuechains (id_status, nama_status) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, valuechain.Id_ValueChain, valuechain.Nama_ValueChain)
	helper.PanicError(err)
	fmt.Println(valuechain)
	return valuechain
}
