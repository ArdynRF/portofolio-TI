package valuechains

type ValuechainCreateRequest struct {
	Id_ValueChain   int64  `json: "id_valuechain"`
	Nama_ValueChain string `json: "nama_valuechain"`
}
