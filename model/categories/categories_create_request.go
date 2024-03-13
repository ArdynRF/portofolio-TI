package categories

type CategoriesCreateRequest struct {
	Id_Category   int64 `json: "id_category"`
	Nama_Category string `json: "nama_category"`
}