package models

import (
	"mobile-backendgo-spp/config"
	"mobile-backendgo-spp/entity"
)

func FindAllTransaksiPeriode(nis, tahunajaran int) ([]*entity.TransaksiLain, error) {
	var result []*entity.TransaksiLain
	tx := config.DB.Table("transaksi_lain").
		Select("transaksi_lain.*, jenis_bayar.*").
		Joins("left join jenis_bayar on transaksi_lain.id_tagihan = jenis_bayar.id_jenisbayar").
		Where("NIS", nis).
		Where("tahun_ajaran", tahunajaran).
		Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}
