package models

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/entity"
)

func FindTahunAjaranActive() (*entity.TahunAjaran, error) {
	var result *entity.TahunAjaran
	tx := config.DB.Where("status", "aktif").First(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}
