package model

import (
	"kreditplus-test/dto"
	"time"
)

type Customer struct {
	ID           string    `db:"id"`
	NIK          string    `db:"nik"`
	FullName     string    ` db:"full_name"`
	LegalName    string    ` db:"legal_name"`
	BirthPlace   string    `db:"tempat_lahir"`
	BirthDate    time.Time `db:"tanggal_lahir"`
	Salary       float64   `db:"gaji"`
	KTPPhoto     string    ` db:"foto_ktp"`
	SelfiePhoto  string    ` db:"foto_selfie"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    int64     `db:"created_at"`
	UpdatedAt    int64     `db:"updated_at"`
	IsDeleted    int64     `db:"is_deleted"`
}

func (c Customer) ToDTO() dto.UserResponse {
	return dto.UserResponse{
		NIK:        c.NIK,
		FullName:   c.FullName,
		LegalName:  c.LegalName,
		BirthPlace: c.BirthPlace,
		BirthDate:  c.BirthDate,
		Salary:     c.Salary,
	}
}
