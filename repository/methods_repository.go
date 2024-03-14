package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type MethodsRepository interface {
	Create(ctx context.Context, tx *sql.Tx, method entity.Methods) entity.Methods
}
