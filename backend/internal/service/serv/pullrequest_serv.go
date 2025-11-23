// Package serv содержит методы для реализации бизнес-логики (пока пусто, так как проект небольшой)
package serv

import (
	"github.com/ArtemST2006/Avito_internship/backend/internal/repository"
	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
)

type PullRequestService struct {
	repo *repository.Repository
}

func NewPullRequestService(repo *repository.Repository) *PullRequestService {
	return &PullRequestService{
		repo: repo,
	}
}

func (p *PullRequestService) MergePR(entitie schemas.PullRqMergeRequest) (schemas.PullRqResponse, error) {
	return p.repo.PullRequest.MergePR(entitie)
}

func (p *PullRequestService) CreatePR(entitie schemas.CreatePullRequestRequest) (schemas.PullRqResponse, error) {
	return p.repo.PullRequest.CreatePR(entitie)
}

func (p *PullRequestService) ChangeAuthorPR(entitie schemas.PRChangeAuthorRequest) (schemas.PRChangeAuthorResponse, error) {
	return p.repo.PullRequest.ChangeAuthorPR(entitie)
}
