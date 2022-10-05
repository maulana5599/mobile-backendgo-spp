package models

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/entity"
)

func FindTransaksiPeriode(nis, tahunajaran int) ([]*entity.TransaksiSpp, error) {
	var transaksi []*entity.TransaksiSpp
	tx := config.DB.
		Where("NIS", nis).
		Where("tahun_ajaran", tahunajaran).
		Find(&transaksi)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return transaksi, nil
}
