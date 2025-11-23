package serv

import (
	"github.com/ArtemST2006/Avito_internship/backend/internal/repository"
	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
)

type StatisticService struct {
	repo *repository.Repository
}

func NewStatisticService(repo *repository.Repository) *StatisticService {
	return &StatisticService{
		repo: repo,
	}
}

func (s *StatisticService) Statistic() (schemas.StatisticResponse, error) {
	return s.repo.Statistic.Statistic()
}
