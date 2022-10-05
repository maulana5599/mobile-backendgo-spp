package models

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/entity"
)

func FindTransaksiDetailPeriode(nis, tahunajaran int) ([]*entity.TransaksiSpp, error) {
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

func FindTransaksiPeriode(nis int) ([]*entity.TransaksiSppPeriode, error) {
	var transaksi []*entity.TransaksiSppPeriode

	tx := config.DB.Table("transaksi").
		Where("NIS", nis).
		Select("transaksi.*, tahun_ajaran.*").
		Joins("left join tahun_ajaran on transaksi.tahun_ajaran = tahun_ajaran.id_tahunajaran").
		Group("transaksi.tahun_ajaran").
		Scan(&transaksi)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return transaksi, nil
}
