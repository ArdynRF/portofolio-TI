package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type ArchitecturesRepositoryImpl struct{}

func NewArchitecturesRepositoryImpl() ArchitecturesRepository {
	return &ArchitecturesRepositoryImpl{}
}
func (repository *ArchitecturesRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, architecture entity.Architectures) entity.Architectures {
	SQL := "INSERT INTO architectures (id_architecture,nama_architecture) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, architecture.Id_Architecture, architecture.Nama_Architecture)
	helper.PanicError(err)
	fmt.Println(architecture)
	return architecture
}
