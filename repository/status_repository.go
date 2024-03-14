package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type StatusRepository interface {
	Create(ctx context.Context, tx *sql.Tx, status entity.Status) entity.Status
}
