package entity

// tabel 13
type Aplikasi struct {
	Id_App            int64
	NamaApp           string
	Tahun             int64
	Alias             string
	URL               string
	BussinesOwner     string
	FrontEndLang      string
	FrontEndFrame     string
	BackEndLang       string
	BackEndFrame      string
	Acquisition       string
	Foto              string
	Id_Perusahaan     int64
	Id_Category       int64
	Id_Status         int64
	Id_ValueChain     int64
	Id_Provider       int64
	Id_Assets         int64
	Id_Contract       int64
	Id_Vendor         int64
	Id_Method         int64
	Id_Type           int64
	Id_Architecture   int64
	Id_Infrastructure int64
}
