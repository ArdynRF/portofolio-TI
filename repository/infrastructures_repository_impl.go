package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type InfrastructuresRepositoryImpl struct{}

func NewInfrastructuresRepositoryImpl() InfrastructuresRepository {
	return &InfrastructuresRepositoryImpl{}
}
func (repository *InfrastructuresRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, infrastructure entity.Infrastructures) entity.Infrastructures {
	SQL := "INSERT INTO infrastructures (id_infrastructure, nama_infrastructure) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, infrastructure.Id_Infrastructure, infrastructure.Nama_Infrastructure)
	helper.PanicError(err)
	fmt.Println(infrastructure)
	return infrastructure
}
