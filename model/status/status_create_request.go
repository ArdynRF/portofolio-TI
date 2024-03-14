package status

type StatusCreateRequest struct {
	Id_Status   int64  `json: "id_status"`
	Nama_Status string `validate:"required" json: "nama_status"`
}
