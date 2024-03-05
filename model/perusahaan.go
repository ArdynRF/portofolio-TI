package model

// tabel 1
type tabel_perusahaan struct {
	ID_Perusahaan string
	Perusahaan    string
}

// tabel 2
type tabel_anakperusahaan struct {
	ID_AnakPerusahaan string
	AnakPerusahaan    string
}

// tabel 3
type tabel_contractstatus struct {
	ID_ContractStat string
	ContractStatus  string
}

// tabel 4
type tabel_status struct {
	ID_Stat string
	status  string
}

// tabel 5
type tabel_assetstatus struct {
	ID_AssetStat string
	AssetStatus  string
}

// tabel 6
type tabel_provider struct {
	ID_Provider string
	Provider    string
}

// tabel 7
type tabel_valuechain struct {
	ID_ValueChain string
	ValueChain    string
}

// tabel 8
type tabel_acquisitioncost struct {
	ID_Acquisition string
	Acquisition    string
}

// tabel 9
type tabel_developmentvendor struct {
	ID_DevVendor string
	DevVendor    string
}

// tabel 10
type tabel_developmentmethod struct {
	ID_DevMethod string
	DevMethod    string
}

// tabel 11
type tabel_developmenttype struct {
	ID_DevType string
	DevType    string
}

// tabel 12
type tabel_architecture struct {
	ID_Architecture string
	Architecture    string
}

// tabel 13
type tabel_infrastructure struct {
	ID_Structure   string
	Infrastructure string
}

// tabel 14
type tabel_aplikasi struct {
	ID_App            string
	NamaApp           string
	Tahun             string
	Alias             string
	URL               string
	BussinesOwner     string
	FrontEndLang      string
	FrontEndFrame     string
	BackEndLang       string
	BackEndFrame      string
	ID_Perusahaan     string
	ID_AnakPerusahaan string
	ID_Stat           string
	ID_ValueChain     string
	ID_Provider       string
	ID_AssetStat      string
	ID_Acquisition    string
	ID_ContractStat   string
	ID_DevVendor      string
	ID_DevMethod      string
	ID_DevType        string
	ID_Architecture   string
	ID_Structure      string
}
