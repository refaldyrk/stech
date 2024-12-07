package dto

import "time"

type LoginCustomerRequest struct {
	NIK      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterCustomerRequest struct {
	NIK        string  `json:"nik" binding:"required"`
	FullName   string  ` json:"full_name" binding:"required"`
	LegalName  string  ` json:"legal_name" binding:"required"`
	BirthPlace string  `json:"tempat_lahir" binding:"required"`
	BirthDate  string  `json:"tanggal_lahir" binding:"required"`
	Salary     float64 `json:"gaji" binding:"required"`
	Password   string  `json:"password" binding:"required"`
}

type UserResponse struct {
	NIK        string    `json:"nik"`
	FullName   string    ` json:"full_name"`
	LegalName  string    ` json:"legal_name"`
	BirthPlace string    `json:"tempat_lahir"`
	BirthDate  time.Time `json:"tanggal_lahir"`
	Salary     float64   `json:"gaji"`
}
