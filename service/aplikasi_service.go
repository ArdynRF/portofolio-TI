package service

import (
	"context"

	"github.com/ArdynRF/Portofolio-TI/model/aplikasi"
)

type AplikasiService interface {
	Create(ctx context.Context, request aplikasi.AplikasiCreateRequest) aplikasi.AplikasiResponse
	Update(ctx context.Context, request aplikasi.AplikasiUpdateRequest) aplikasi.AplikasiResponse
	Delete(ctx context.Context, appId int64)
	FindById(ctx context.Context, appId int64) aplikasi.AplikasiResponse
}
