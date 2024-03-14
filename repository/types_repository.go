package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type TypesRepository interface {
	Create(ctx context.Context, tx *sql.Tx, tipe entity.Types) entity.Types
}
