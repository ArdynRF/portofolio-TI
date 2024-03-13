package perusahaan

type PerusahaanCreateRequest struct {
	Id_Perusahaan   int64 `json: "id_perusahaan"`
	Nama_Perusahaan string `json: "nama_perusahaan"`
}
