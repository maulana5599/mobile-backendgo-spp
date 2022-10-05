package entity

type Siswa struct {
	IdSiswa   int    `json:"id_siswa"`
	Nis       int    `json:"nis"`
	NamaSiswa string `json:"nama_siswa"`
	Jurusan   string `json:"jurusan"`
	Kelas     string `json:"kelas"`
	Status    string `json:"status"`
}
