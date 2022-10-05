package models

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/entity"
)

func FindSiswaByNis(nis int) (*entity.Siswa, error) {
	var siswa *entity.Siswa
	tx := config.DB.Where("NIS", nis).First(&siswa)
	if tx.Error != nil {
		return siswa, tx.Error
	}

	return siswa, nil
}
