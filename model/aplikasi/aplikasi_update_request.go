package aplikasi

type AplikasiUpdateRequest struct {
	Id_App            int64  `json:"id_app"`
	NamaApp           string `json: "nama_app"`
	Tahun             int64  `json: "tahun"`
	Alias             string `json: "alias"`
	URL               string `json: "url"`
	BussinesOwner     string `json: "bussines_owner"`
	FrontEndLang      string `json: "fe_language"`
	FrontEndFrame     string `json: "fe_frame"`
	BackEndLang       string `json: "be_languge"`
	BackEndFrame      string `json: "be_frame"`
	Acquisition       string `json: "acquisition"`
	Foto              string `json: "foto"`
	Provider          string `json: "provider"`
	Vendor            string `json: "vendor"`
	Id_Perusahaan     int64  `json: "id_perusahaan"`
	Id_Category       int64  `json: "id_category"`
	Id_Status         int64  `json: "id_status"`
	Id_ValueChain     int64  `json: "id_valuechain"`
	Id_Assets         int64  `json: "id_assets"`
	Id_Contract       int64  `json: "id_contract"`
	Id_Method         int64  `json: "id_method"`
	Id_Type           int64  `json: "id_type"`
	Id_Architecture   int64  `json: "id_architecture"`
	Id_Infrastructure int64  `json: "id_infrastructure"`
	Updated_At        int64  `json: updated_at`
	Review            bool   `json: review`
	Approval          bool   `json: approval`
}
