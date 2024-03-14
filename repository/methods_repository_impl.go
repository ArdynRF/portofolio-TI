package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type MethodsRepositoryImpl struct{}

func NewMethodsRepositoryImpl() MethodsRepository {
	return &MethodsRepositoryImpl{}
}
func (repository *MethodsRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, method entity.Methods) entity.Methods {
	SQL := "INSERT INTO methods (id_method, nama_method) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, method.Id_Method, method.Nama_Method)
	helper.PanicError(err)
	fmt.Println(method)
	return method
}
