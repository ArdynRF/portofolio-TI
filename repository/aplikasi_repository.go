package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type AplikasiRepository interface {
	Create(ctx context.Context, tx *sql.Tx, aplikasi entity.Aplikasi) entity.Aplikasi
}
