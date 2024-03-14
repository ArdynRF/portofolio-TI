package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type ContractsRepositoryImpl struct{}

func NewContractsRepositoryImpl() ContractsRepository {
	return &ContractsRepositoryImpl{}
}
func (repository *ContractsRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, contract entity.Contracts) entity.Contracts {
	SQL := "INSERT INTO contracts (id_contract, nama_contract) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, contract.Id_Contract, contract.Nama_Contract)
	helper.PanicError(err)
	fmt.Println(contract)
	return contract
}
