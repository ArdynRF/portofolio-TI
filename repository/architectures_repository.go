package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type ArchitecturesRepository interface {
	Create(ctx context.Context, tx *sql.Tx, architecture entity.Architectures) entity.Architectures
}
