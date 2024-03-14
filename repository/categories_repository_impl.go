package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type CategoriesRepositoryImpl struct{}

func NewCategoriesRepositoryImpl() CategoriesRepository {
	return &CategoriesRepositoryImpl{}
}
func (repository *CategoriesRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category entity.Categories) entity.Categories {
	SQL := "INSERT INTO categories (id_category, nama_category) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, category.Id_Category, category.Nama_Category)
	helper.PanicError(err)
	fmt.Println(category)
	return category
}
