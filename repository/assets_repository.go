package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type AssetsRepository interface {
	Create(ctx context.Context, tx *sql.Tx, asset entity.Assets) entity.Assets
}
