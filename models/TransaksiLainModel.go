package models

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/entity"
)

func FindAllTransaksiPeriode(nis, tahunajaran int) ([]*entity.TransaksiLain, error) {
	var result []*entity.TransaksiLain
	tx := config.DB.Where("NIS", nis).
		Where("tahun_ajaran", tahunajaran).
		Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}
