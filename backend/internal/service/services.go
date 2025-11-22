package service

import (
	"github.com/ArtemST2006/Avito_internship/backend/internal/repository"
	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
	"github.com/ArtemST2006/Avito_internship/backend/internal/service/serv"
)

type PullRequest interface {
	MergePR(schemas.PullRqMergeRequest) (schemas.PullRqResponse, error)
	CreatePR(schemas.CreatePullRequestRequest) (schemas.PullRqResponse, error)
	ChangeAuthorPR(schemas.PRChangeAuthorRequest) (schemas.PRChangeAuthorResponse, error)
}

type Team interface {
	GetTeam(string) (schemas.GetTeamResponse, error)
	AddTeam(schemas.CreateTeamRequest) (schemas.CreateTeamResponse, error)
}

type User interface {
	SetIsActive(schemas.ActieveUserRequest) (schemas.ActiveUserResponse, error)
	GetUserReview(string) (schemas.GetUserPRResponse, error)
}

type Service struct {
	Team
	PullRequest
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		PullRequest: serv.NewPullRequestService(repo),
		Team:        serv.NewTeamService(repo),
		User:        serv.NewUserService(repo),
	}
}
