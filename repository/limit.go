package repository

import (
	"database/sql"
	"fmt"
	"kreditplus-test/dto"
	"kreditplus-test/model"
)

type LimitRepository struct {
	db *sql.DB
}

func NewLimitRepository(db *sql.DB) *LimitRepository {
	return &LimitRepository{db}
}

func (r *LimitRepository) Insert(limit model.Limit) error {
	query := `INSERT INTO LimitCustomer (id, konsumen_id, tenor, limit_pinjaman) 
	VALUES (?, ?, ?, ?)`

	_, err := r.db.Exec(query, limit.ID, limit.KonsumenID, limit.Tenor, limit.LimitPinjaman)
	return err
}

func (r *LimitRepository) FindString(key, val string) (model.Limit, error) {
	var limit model.Limit
	query := fmt.Sprintf(`
		SELECT l.id, l.konsumen_id, l.tenor, l.limit_pinjaman, 
		       k.id, k.nik, k.full_name, k.legal_name, k.tempat_lahir, k.tanggal_lahir, k.gaji, k.foto_ktp, k.foto_selfie
		FROM LimitCustomer l
		JOIN Konsumen k ON l.konsumen_id = k.id
		WHERE %s = ?
	`, key)

	row := r.db.QueryRow(query, val)
	err := row.Scan(
		&limit.ID, &limit.KonsumenID, &limit.Tenor, &limit.LimitPinjaman,
		&limit.Konsumen.ID, &limit.Konsumen.NIK, &limit.Konsumen.FullName, &limit.Konsumen.LegalName, &limit.Konsumen.BirthPlace,
		&limit.Konsumen.BirthDate, &limit.Konsumen.Salary, &limit.Konsumen.KTPPhoto, &limit.Konsumen.SelfiePhoto,
	)

	if err != nil {
		return model.Limit{}, err
	}

	return limit, nil
}

func (r *LimitRepository) FindTenor(id string, tenor int) (dto.LimitResponse, error) {
	var limit model.Limit
	query := fmt.Sprintf(`
		SELECT l.id, l.konsumen_id, l.tenor, l.limit_pinjaman, 
		       k.id, k.nik, k.full_name, k.legal_name, k.tempat_lahir, k.tanggal_lahir, k.gaji, k.foto_ktp, k.foto_selfie
		FROM LimitCustomer l
		JOIN Konsumen k ON l.konsumen_id = k.id
		WHERE l.konsumen_id = ? AND l.tenor = ?`)

	row := r.db.QueryRow(query, id, tenor)
	err := row.Scan(
		&limit.ID, &limit.KonsumenID, &limit.Tenor, &limit.LimitPinjaman,
		&limit.Konsumen.ID, &limit.Konsumen.NIK, &limit.Konsumen.FullName, &limit.Konsumen.LegalName, &limit.Konsumen.BirthPlace,
		&limit.Konsumen.BirthDate, &limit.Konsumen.Salary, &limit.Konsumen.KTPPhoto, &limit.Konsumen.SelfiePhoto,
	)

	if err != nil {
		return dto.LimitResponse{}, err
	}

	return limit.ToDTO(), nil
}

func (r *LimitRepository) FindAll(konsumenID string, page, limit int) ([]dto.LimitResponse, error) {
	var limits []model.Limit
	offset := (page - 1) * limit

	query := `
		SELECT l.id, l.konsumen_id, l.tenor, l.limit_pinjaman, 
		       k.id, k.nik, k.full_name, k.legal_name, k.tempat_lahir, k.tanggal_lahir, k.gaji, k.foto_ktp, k.foto_selfie
		FROM LimitCustomer l
		JOIN Konsumen k ON l.konsumen_id = k.id
		WHERE l.konsumen_id = ?
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, konsumenID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var limit model.Limit
		if err := rows.Scan(
			&limit.ID, &limit.KonsumenID, &limit.Tenor, &limit.LimitPinjaman,
			&limit.Konsumen.ID, &limit.Konsumen.NIK, &limit.Konsumen.FullName, &limit.Konsumen.LegalName, &limit.Konsumen.BirthPlace,
			&limit.Konsumen.BirthDate, &limit.Konsumen.Salary, &limit.Konsumen.KTPPhoto, &limit.Konsumen.SelfiePhoto,
		); err != nil {
			return nil, err
		}
		limits = append(limits, limit)
	}

	var response []dto.LimitResponse
	for i := range limits {
		response = append(response, limits[i].ToDTO())
	}
	return response, nil
}

func (r *LimitRepository) Update(limit model.Limit) error {
	query := `
		UPDATE LimitCustomer 
		SET tenor = ?, limit_pinjaman = ? 
		WHERE id = ?
	`

	_, err := r.db.Exec(query, limit.Tenor, limit.LimitPinjaman, limit.ID)
	return err
}

func (r *LimitRepository) Delete(id string) error {
	query := `DELETE FROM Limit WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
