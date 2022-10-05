package entity

import "time"

type TransaksiSpp struct {
	IdTransaksi  int
	NIS          int
	TahunAjaran  int
	Kelas        int
	Jurusan      int
	JumlahBayar  int
	TanggalBayar time.Time
	BulanBayar   string
}

func (TransaksiSpp) TableName() string {
	return "transaksi"
}
