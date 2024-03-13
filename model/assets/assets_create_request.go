package assets

type AssetsCreateRequest struct {
	Id_Asset int64 `json: "id_asset"`
	Nama_Asset   string `json: "nama_asset"`
}