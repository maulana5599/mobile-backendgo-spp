package models

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/entity"
)

func FindSiswaByNis(nis int) (*entity.Siswa, error) {
	var siswa *entity.Siswa
	tx := config.DB.Where("NIS", nis).
		Select("siswa.*, jurusan.jurusan").
		Joins("join jurusan on siswa.jurusan = jurusan.id_jurusan").
		First(&siswa)
	if tx.Error != nil {
		return siswa, tx.Error
	}

	return siswa, nil
}
