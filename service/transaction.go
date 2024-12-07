package service

import (
	"fmt"
	"github.com/rs/xid"
	"kreditplus-test/dto"
	"kreditplus-test/model"
	"kreditplus-test/repository"
	"time"
)

type TransaksiService struct {
	transaksiRepo *repository.TransaksiRepository
	limitRepo     *repository.LimitRepository
}

func NewTransaksiService(transaksiRepo *repository.TransaksiRepository, limitRepo *repository.LimitRepository) *TransaksiService {
	return &TransaksiService{
		transaksiRepo: transaksiRepo,
		limitRepo:     limitRepo,
	}
}

func (s TransaksiService) CreateTransaksi(userid string, req dto.TransaksiRequestDTO) (dto.TransaksiResponseDTO, error) {
	limits, err := s.limitRepo.FindAll(userid, 1, 100)
	if err != nil {
		return dto.TransaksiResponseDTO{}, fmt.Errorf("error fetching limit data: %v", err)
	}

	totalCicilan := req.JumlahCicilan + req.AdminFee

	var limitId string
	var limitPinjaman float64
	for _, limit := range limits {
		if limit.LimitPinjaman >= totalCicilan {
			limitPinjaman = limit.LimitPinjaman
			limitId = limit.ID
		}
	}

	if limitPinjaman == 0 {
		return dto.TransaksiResponseDTO{}, fmt.Errorf("no sufficient limit found for the given tenor")
	}

	if totalCicilan > limitPinjaman {
		return dto.TransaksiResponseDTO{}, fmt.Errorf("requested transaction amount exceeds available limit")
	}

	transaksi := model.Transaksi{
		ID:            xid.New().String(),
		KonsumenID:    userid,
		NomorKontrak:  req.NomorKontrak,
		NamaAsset:     req.NamaAsset,
		OTR:           req.OTR,
		AdminFee:      req.AdminFee,
		JumlahCicilan: totalCicilan,
		JumlahBunga:   req.JumlahBunga,
		CreatedAt:     time.Now().Unix(),
	}

	err = s.transaksiRepo.Insert(transaksi)
	if err != nil {
		return dto.TransaksiResponseDTO{}, fmt.Errorf("error inserting transaksi: %v", err)
	}

	//Update Limit
	limitResponse, err := s.limitRepo.FindString("l.id", limitId)
	if err != nil {
		return dto.TransaksiResponseDTO{}, err
	}

	limitResponse.LimitPinjaman = limitResponse.LimitPinjaman - totalCicilan

	err = s.limitRepo.Update(limitResponse)
	if err != nil {
		return dto.TransaksiResponseDTO{}, err
	}

	return dto.TransaksiResponseDTO{
		TransactionID:  transaksi.ID,
		Tenor:          limitResponse.Tenor,
		LimitAvailable: limitResponse.LimitPinjaman,
	}, nil
}

func (s TransaksiService) GetAllTransaction(userID string, page, limit int) ([]dto.TransaksiResponse, error) {
	var response []dto.TransaksiResponse

	transaksi, err := s.transaksiRepo.FindAll(userID, page, limit)
	if err != nil {
		return nil, err
	}

	for i := range transaksi {
		response = append(response, transaksi[i].ToDto())
	}

	return response, nil
}
