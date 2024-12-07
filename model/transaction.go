package model

import "kreditplus-test/dto"

type Transaksi struct {
	ID            string  `db:"id"`
	KonsumenID    string  `db:"konsumen_id"`
	NomorKontrak  string  `db:"nomor_kontrak"`
	NamaAsset     string  `db:"nama_asset"`
	OTR           float64 `db:"otr"`
	AdminFee      float64 `db:"admin_fee"`
	JumlahCicilan float64 `db:"jumlah_cicilan"`
	JumlahBunga   float64 `db:"jumlah_bunga"`
	CreatedAt     int64   `db:"created_at"`
}

func (t Transaksi) ToDto() dto.TransaksiResponse {
	return dto.TransaksiResponse{
		ID:            t.ID,
		KonsumenID:    t.KonsumenID,
		NomorKontrak:  t.NomorKontrak,
		NamaAsset:     t.NamaAsset,
		OTR:           t.OTR,
		AdminFee:      t.AdminFee,
		JumlahCicilan: t.JumlahCicilan,
		JumlahBunga:   t.JumlahBunga,
		CreatedAt:     t.CreatedAt,
	}
}
