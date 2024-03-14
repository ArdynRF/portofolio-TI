package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ArdynRF/Portofolio-TI/helper"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
)

type AssetsRepositoryImpl struct{}

func NewAssetsRepositoryImpl() AssetsRepository {
	return &AssetsRepositoryImpl{}
}
func (repository *AssetsRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, asset entity.Assets) entity.Assets {
	SQL := "INSERT INTO assets (id_asset, nama_asset) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, SQL, asset.Id_Asset, asset.Nama_Asset)
	helper.PanicError(err)
	fmt.Println(asset)
	return asset
}
