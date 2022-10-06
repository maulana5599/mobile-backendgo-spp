package models

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/entity"
)

/**
Get data pembayaran per jurusan
*/
func FindTunggakanSiswa(jurusan, kelas, tahunajaran int) (*entity.Tunggakan, error) {
	var result *entity.Tunggakan
	tx := config.DB.Table("pembayaran").
		Where("jurusan", jurusan).
		Where("kelas", kelas).
		Where("tahun_ajaran", tahunajaran).
		First(&result)

	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

func FindTunggakanLainSiswa(jurusan, kelas, tahunajaran int) ([]*entity.TunggakanLain, error) {
	var result []*entity.TunggakanLain
	tx := config.DB.Raw("SELECT * FROM pembayaran_lainnya_view where jurusan = ? and kelas = ? and tahun_ajaran = ?", jurusan, kelas, tahunajaran).Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}
