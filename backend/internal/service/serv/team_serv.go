package serv

import (
	"github.com/ArtemST2006/Avito_internship/backend/internal/repository"
	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
)

type TeamService struct {
	repo *repository.Repository
}

func NewTeamService(repo *repository.Repository) *TeamService {
	return &TeamService{
		repo: repo,
	}
}

func (u *TeamService) GetTeam(tm string) (schemas.GetTeamResponse, error) {
	return u.repo.Team.GetTeam(tm)
}

func (u *TeamService) AddTeam(req schemas.CreateTeamRequest) (schemas.CreateTeamResponse, error) {
	return u.repo.AddTeam(req)
}
