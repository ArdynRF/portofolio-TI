package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type PerusahaanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, perusahaan entity.Perusahaan) entity.Perusahaan
}
