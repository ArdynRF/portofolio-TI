package vendors

type VendorsCreateRequest struct {
	Id_Vendor   int64  `json: "id_vendor"`
	Nama_Vendor string `json: "nama_vendor"`
}
