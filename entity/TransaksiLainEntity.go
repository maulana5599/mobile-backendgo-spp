package entity

import "time"

type TransaksiLain struct {
	IdTransaksi  int
	NIS          int
	TahunAjaran  int
	Kelas        int
	Jurusan      int
	JumlahBayar  int
	TanggalBayar time.Time
	BulanBayar   string
}

func (TransaksiLain) TableName() string {
	return "transaksi_lain"
}
