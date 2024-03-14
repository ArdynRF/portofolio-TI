package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type ValuechainRepository interface {
	Create(ctx context.Context, tx *sql.Tx, tipe entity.Valuechain) entity.Valuechain
}
