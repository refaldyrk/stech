package service

import (
	"kreditplus-test/dto"
	"kreditplus-test/repository"
)

type LimitService struct {
	limitRepository *repository.LimitRepository
}

func NewLimitService(limitRepository *repository.LimitRepository) *LimitService {
	return &LimitService{limitRepository}
}

func (l LimitService) GetAllLimitByUserID(userID string, page, limit int) ([]dto.LimitResponse, error) {
	limits, err := l.limitRepository.FindAll(userID, page, limit)
	if err != nil {
		return nil, err
	}

	return limits, err
}

func (l LimitService) GetLimitByTenor(id string, tenor int) (dto.LimitResponse, error) {
	return l.limitRepository.FindTenor(id, tenor)
}
