package postgres

import (
	"encoding/json"
	"errors"

	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
	projerrors "github.com/ArtemST2006/Avito_internship/backend/pkg/errors"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (u *UserPostgres) SetIsActive(req schemas.ActieveUserRequest) (schemas.ActiveUserResponse, error) {
	response := schemas.ActiveUserResponse{}

	var user schemas.User
	err := u.db.Where("user_id = ?", req.UserID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, projerrors.ErrNotFound
		}
		return response, err
	}

	user.IsActive = req.IsActive
	err = u.db.Save(&user).Error
	if err != nil {
		return response, err
	}

	response.User = schemas.User{
		UserId:   user.UserId,
		UserName: user.UserName,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}

	return response, nil
}

func (u *UserPostgres) GetUserReview(userId string) (schemas.GetUserPRResponse, error) {
	response := schemas.GetUserPRResponse{}

	var user schemas.User
	err := u.db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, projerrors.ErrNotFound
		}
		return response, err
	}

	jsonQuery, err := json.Marshal([]string{userId})
	if err != nil {
		return response, err
	}

	var pullRequests []schemas.PullRequest
	err = u.db.Raw("SELECT * FROM pull_requests WHERE assigned_reviewers @> ?", string(jsonQuery)).Scan(&pullRequests).Error
	if err != nil {
		return response, err
	}

	prInfoList := make([]schemas.PullRequestInfo, 0, len(pullRequests))
	for _, pr := range pullRequests {
		prInfo := schemas.PullRequestInfo{
			PullRequestID:   pr.PullRequestID,
			PullRequestName: pr.PullRequestName,
			AuthorID:        pr.AuthorID,
			Status:          pr.Status,
		}
		prInfoList = append(prInfoList, prInfo)
	}

	response.UserID = userId
	response.PullRequests = prInfoList

	return response, nil
}
