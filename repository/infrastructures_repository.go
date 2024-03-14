package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type InfrastructuresRepository interface {
	Create(ctx context.Context, tx *sql.Tx, infrastructure entity.Infrastructures) entity.Infrastructures
}
