package repository

import (
	"database/sql"
	"fmt"
	"kreditplus-test/model"
	"time"
)

type CustomerRepository struct {
	db              *sql.DB
	LimitRepository *LimitRepository
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db, NewLimitRepository(db)}
}

func (r *CustomerRepository) Insert(customer model.Customer) error {
	query := `
		INSERT INTO Konsumen (id, nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie, password_hash, created_at, updated_at, is_deleted) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query, customer.ID, customer.NIK, customer.FullName, customer.LegalName, customer.BirthPlace, customer.BirthDate, customer.Salary, customer.KTPPhoto, customer.SelfiePhoto, customer.PasswordHash, customer.CreatedAt, customer.UpdatedAt, customer.IsDeleted)
	return err
}

func (r *CustomerRepository) Find(key, val string) (model.Customer, error) {
	var customer model.Customer
	query := fmt.Sprintf(`SELECT id, nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie, password_hash, created_at, updated_at, is_deleted FROM Konsumen WHERE %s = ?`, key)

	row := r.db.QueryRow(query, val)
	err := row.Scan(&customer.ID, &customer.NIK, &customer.FullName, &customer.LegalName, &customer.BirthPlace, &customer.BirthDate, &customer.Salary, &customer.KTPPhoto, &customer.SelfiePhoto, &customer.PasswordHash, &customer.CreatedAt, &customer.UpdatedAt, &customer.IsDeleted)

	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

func (r *CustomerRepository) FindAll(page, pageSize int) ([]model.Customer, error) {
	var customers []model.Customer
	offset := (page - 1) * pageSize
	query := `SELECT id, nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie, password_hash, created_at, updated_at, is_deleted FROM Konsumen LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.ID, &customer.NIK, &customer.FullName, &customer.LegalName, &customer.BirthPlace, &customer.BirthDate, &customer.Salary, &customer.KTPPhoto, &customer.SelfiePhoto, &customer.PasswordHash, &customer.CreatedAt, &customer.UpdatedAt, &customer.IsDeleted); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *CustomerRepository) Update(customer model.Customer) error {
	query := `
		UPDATE Konsumen 
		SET nik = ?, full_name = ?, legal_name = ?, tempat_lahir = ?, tanggal_lahir = ?, gaji = ?, foto_ktp = ?, foto_selfie = ?, password_hash = ?, updated_at = ?, is_deleted = ? 
		WHERE id = ?
	`

	_, err := r.db.Exec(query, customer.NIK, customer.FullName, customer.LegalName, customer.BirthPlace, customer.BirthDate, customer.Salary, customer.KTPPhoto, customer.SelfiePhoto, customer.PasswordHash, time.Now().Unix(), customer.IsDeleted, customer.ID)
	return err
}

func (r *CustomerRepository) Delete(id string) error {
	query := `UPDATE Konsumen SET is_deleted = ? WHERE id = ?`
	_, err := r.db.Exec(query, true, id)
	return err
}
