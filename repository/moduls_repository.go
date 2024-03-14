package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type ModulsRepository interface {
	Create(ctx context.Context, tx *sql.Tx, moduls entity.Moduls) entity.Moduls
}
