package utils

import (
	"github.com/ArdynRF/Portofolio-TI/model/aplikasi"
	"github.com/ArdynRF/Portofolio-TI/model/entity"
	"github.com/ArdynRF/Portofolio-TI/model/web"
)

func UserResponse(user entity.Users) web.UsersResponse {
	return web.UsersResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Nama:      user.Nama,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func AplikasiResponse(app entity.Aplikasi) aplikasi.AplikasiResponse {
	return aplikasi.AplikasiResponse{
		Id_App:            app.Id_App,
		NamaApp:           app.NamaApp,
		Tahun:             app.Tahun,
		Alias:             app.Alias,
		URL:               app.URL,
		BussinesOwner:     app.BussinesOwner,
		FrontEndLang:      app.FrontEndLang,
		FrontEndFrame:     app.FrontEndFrame,
		BackEndLang:       app.BackEndLang,
		BackEndFrame:      app.BackEndFrame,
		Acquisition:       app.Acquisition,
		Foto:              app.Foto,
		Provider:          app.Provider,
		Vendor:            app.Vendor,
		Id_Perusahaan:     app.Id_Perusahaan,
		Id_Category:       app.Id_Category,
		Id_Status:         app.Id_Status,
		Id_ValueChain:     app.Id_ValueChain,
		Id_Assets:         app.Id_Assets,
		Id_Contract:       app.Id_Contract,
		Id_Method:         app.Id_Method,
		Id_Type:           app.Id_Type,
		Id_Architecture:   app.Id_Architecture,
		Id_Infrastructure: app.Id_Infrastructure,
		Created_At:        app.Created_At,
		Updated_At:        app.Updated_At,
	}
}
