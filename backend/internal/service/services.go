package service

import "github.com/ArtemST2006/Avito_internship/backend/internal/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
