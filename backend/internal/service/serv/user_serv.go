package serv

import (
	"github.com/ArtemST2006/Avito_internship/backend/internal/repository"
	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) SetIsActive(req schemas.ActieveUserRequest) (schemas.ActiveUserResponse, error) {
	return u.repo.User.SetIsActive(req)
}

func (u *UserService) GetUserReview(ui string) (schemas.GetUserPRResponse, error) {
	return u.repo.User.GetUserReview(ui)
}
