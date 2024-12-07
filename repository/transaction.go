package repository

import (
	"database/sql"
	"fmt"
	"kreditplus-test/model"
)

type TransaksiRepository struct {
	db *sql.DB
}

func NewTransaksiRepository(db *sql.DB) *TransaksiRepository {
	return &TransaksiRepository{db}
}

func (r *TransaksiRepository) Insert(transaksi model.Transaksi) error {
	query := `
		INSERT INTO Transaksi (id, konsumen_id, nomor_kontrak, nama_asset, otr, admin_fee, jumlah_cicilan, jumlah_bunga, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, transaksi.ID, transaksi.KonsumenID, transaksi.NomorKontrak, transaksi.NamaAsset, transaksi.OTR, transaksi.AdminFee, transaksi.JumlahCicilan, transaksi.JumlahBunga, transaksi.CreatedAt)
	return err
}

func (r *TransaksiRepository) Find(key, val string) (model.Transaksi, error) {
	var transaksi model.Transaksi
	query := fmt.Sprintf(`
		SELECT id, konsumen_id, nomor_kontrak, nama_asset, otr, admin_fee, jumlah_cicilan, jumlah_bunga, created_at
		FROM Transaksi
		WHERE %s = ?
	`, key)

	row := r.db.QueryRow(query, val)
	err := row.Scan(&transaksi.ID, &transaksi.KonsumenID, &transaksi.NomorKontrak, &transaksi.NamaAsset, &transaksi.OTR, &transaksi.AdminFee, &transaksi.JumlahCicilan, &transaksi.JumlahBunga, &transaksi.CreatedAt)

	if err != nil {
		return model.Transaksi{}, err
	}

	return transaksi, nil
}

func (r *TransaksiRepository) FindAll(userId string, page, pageSize int) ([]model.Transaksi, error) {
	var transaksiList []model.Transaksi
	offset := (page - 1) * pageSize
	query := `
		SELECT id, konsumen_id, nomor_kontrak, nama_asset, otr, admin_fee, jumlah_cicilan, jumlah_bunga, created_at
		FROM Transaksi WHERE konsumen_id = ?
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, userId, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaksi model.Transaksi
		if err := rows.Scan(&transaksi.ID, &transaksi.KonsumenID, &transaksi.NomorKontrak, &transaksi.NamaAsset, &transaksi.OTR, &transaksi.AdminFee, &transaksi.JumlahCicilan, &transaksi.JumlahBunga, &transaksi.CreatedAt); err != nil {
			return nil, err
		}
		transaksiList = append(transaksiList, transaksi)
	}

	return transaksiList, nil
}

func (r *TransaksiRepository) Update(transaksi model.Transaksi) error {
	query := `
		UPDATE Transaksi
		SET konsumen_id = ?, nomor_kontrak = ?, nama_asset = ?, otr = ?, admin_fee = ?, jumlah_cicilan = ?, jumlah_bunga = ?, created_at = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(query, transaksi.KonsumenID, transaksi.NomorKontrak, transaksi.NamaAsset, transaksi.OTR, transaksi.AdminFee, transaksi.JumlahCicilan, transaksi.JumlahBunga, transaksi.CreatedAt, transaksi.ID)
	return err
}

func (r *TransaksiRepository) Delete(id string) error {
	query := `DELETE FROM Transaksi WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
