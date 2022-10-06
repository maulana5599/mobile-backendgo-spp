package entity

type Tunggakan struct {
	IdSetbayar     int    `json:"id_setbayar"`
	NamaPembayaran string `json:"nama_pembayaran"`
	Kelas          int    `json:"kelas"`
	Jurusan        int    `json:"jurusan"`
	Januari        int    `json:"januari"`
	Februari       int    `json:"februari"`
	Maret          int    `json:"maret"`
	April          int    `json:"april"`
	Mei            int    `json:"mei"`
	Juni           int    `json:"juni"`
	Juli           int    `json:"juli"`
	Agustus        int    `json:"agustus"`
	September      int    `json:"september"`
	Oktober        int    `json:"oktober"`
	November       int    `json:"november"`
	Desember       int    `json:"desember"`
}

type TunggakanLain struct {
	IdPemlainnya    int    `json:"id_pemlainnya"`
	IdTagihan       string `json:"id_tagihan"`
	JenisPembayaran string `json:"jenis_pembayaran"`
	Jurusan         int    `json:"jurusan"`
	Kelas           int    `json:"kelas"`
	TahunAjaran     int    `json:"tahun_ajaran"`
	Angkatan        int    `json:"angkatan"`
	JumlahBayar     int    `json:"jumlah_bayar"`
	JumlahCicilan   int    `json:"jumlah_cicilan"`
}
