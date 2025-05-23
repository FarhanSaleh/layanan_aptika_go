package domain

type PermintaanCountResponse struct {
	Bulan     string `json:"bulan,omitempty"`
	Total     int `json:"total"`
	Diproses  int `json:"diproses"`
	Disetujui int `json:"disetujui"`
	Ditolak   int `json:"ditolak"`
}