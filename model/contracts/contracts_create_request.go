package contracts

type ContractCreatRequest struct {
	Id_Contract int64 `json: "id_contract"`
	Nama_Contract   string `json" "nama_contract"`
}