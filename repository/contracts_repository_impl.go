package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type ContractsRepository interface {
	Create(ctx context.Context, tx *sql.Tx, contract entity.Contracts) entity.Contracts
}
