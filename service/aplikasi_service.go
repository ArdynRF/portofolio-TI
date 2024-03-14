package service

import (
	"context"

	"github.com/ArdynRF/Portofolio-TI/model/aplikasi"
)

type AplikasiService interface {
	Create(ctx context.Context, request aplikasi.AplikasiCreateRequest) aplikasi.AplikasiResponse
	// Update(ctx context.Context, request web.UsersUpdateRequest) web.UsersResponse
	// Delete(ctx context.Context, userId string)
}
