package repository

import (
	"context"
	"database/sql"

	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type CategoriesRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category entity.Categories) entity.Categories
}