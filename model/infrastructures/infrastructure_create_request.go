package infrastructures

type InfrastructuresCreaterRequest struct {
	Id_Infrastructure   int64 `json: "id_infrastructure"`
	Nama_Infrastructure string `json: "nama_infrastructure"`
}
