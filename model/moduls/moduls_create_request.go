package moduls

type ModulsCreateRequest struct {
	Id_Modul    int64  `json: "id_modul"`
	Nama_Modul  string `json: "nama_modul"`
	Id_aplikasi int64  `json: "id_aplikasi"`
}
