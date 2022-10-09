package entity

type Siswa struct {
	IdSiswa     int    `json:"id_siswa"`
	NIS         int    `json:"nis"`
	NamaSiswa   string `json:"nama_siswa"`
	Jurusan     string `json:"jurusan"`
	Kelas       string `json:"kelas"`
	Status      string `json:"status"`
	TempatLahir string `json:"tempat_lahir"`
	Alamat      string `json:"alamat"`
	Telepon     string `json:"telepon"`
}

func (Siswa) TableName() string {
	return "siswa"
}
