package models

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/entity"
)

func FindAllJurusan() ([]*entity.Jurusan, error) {
	var result []*entity.Jurusan
	tx := config.DB.Find(&result)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}
