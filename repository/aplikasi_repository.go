package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type AplikasiRepository interface {
	Create(ctx context.Context, tx *sql.Tx, aplikasi entity.Aplikasi) entity.Aplikasi
	Delete(ctx context.Context, tx *sql.Tx, aplikasi entity.Aplikasi)
	FindById(ctx context.Context, tx *sql.Tx, appId int64) (entity.Aplikasi, error)
	Update(ctx context.Context, tx *sql.Tx, aplikasi entity.Aplikasi) entity.Aplikasi
}
