package entity

import "time"

type TransaksiLain struct {
	IdTransaksi     int       `json:"id_transaksi"`
	NIS             int       `json:"nis"`
	TahunAjaran     int       `json:"tahun_ajaran"`
	Kelas           int       `json:"kelas"`
	Jurusan         int       `json:"jurusan"`
	JumlahBayar     int       `json:"jumlah_bayar"`
	TanggalBayar    time.Time `json:"tanggal_bayar"`
	BulanBayar      string    `json:"bulan_bayar"`
	JenisPembayaran string    `json:"jenis_pembayaran"`
}
