package repository

import (
	"github.com/ArtemST2006/Avito_internship/backend/internal/repository/postgres"
	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
	"gorm.io/gorm"
)

type Team interface {
	GetTeam(string) (schemas.GetTeamResponse, error)
	AddTeam(schemas.CreateTeamRequest) (schemas.CreateTeamResponse, error)
}

type PullRequest interface {
	MergePR(schemas.PullRqMergeRequest) (schemas.PullRqResponse, error)
	CreatePR(schemas.CreatePullRequestRequest) (schemas.PullRqResponse, error)
	ChangeAuthorPR(schemas.PRChangeAuthorRequest) (schemas.PRChangeAuthorResponse, error)
}

type User interface {
	SetIsActive(schemas.ActieveUserRequest) (schemas.ActiveUserResponse, error)
	GetUserReview(string) (schemas.GetUserPRResponse, error)
}

type Repository struct {
	Team
	PullRequest
	User
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		PullRequest: postgres.NewPullRequestRepo(db),
		Team:        postgres.NewTeamRepo(db),
		User:        postgres.NewUserPostgres(db),
	}
}
