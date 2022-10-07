package entity

import "time"

type TransaksiSpp struct {
	IdTransaksi  int       `json:"id_transaksi"`
	NIS          int       `json:"NIS"`
	TahunAjaran  int       `json:"tahun_ajaran"`
	Kelas        int       `json:"kelas"`
	Jurusan      int       `json:"jurusan"`
	JumlahBayar  int       `json:"jumlah_bayar"`
	TanggalBayar time.Time `json:"tanggal_bayar"`
	BulanBayar   string    `json:"bulan_bayar"`
}

type TransaksiSppPeriode struct {
	IdTransaksi  int       `json:"id_transaksi"`
	Nis          int       `json:"NIS"`
	TahunAjaran  int       `json:"tahun_ajaran"`
	Kelas        int       `json:"kelas"`
	Jurusan      int       `json:"jurusan"`
	JumlahBayar  int       `json:"jumlah_bayar"`
	TanggalBayar time.Time `json:"tanggal_bayar"`
	BulanBayar   string    `json:"bulan_bayar"`
	TAjar        string    `json:"t_ajar"`
}

func (TransaksiSpp) TableName() string {
	return "transaksi"
}
