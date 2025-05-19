package domain

type PermintaanCountResponse struct {
	Total     int `json:"total"`
	Diproses  int `json:"diproses"`
	Disetujui int `json:"disetujui"`
	Ditolak   int `json:"ditolak"`
}