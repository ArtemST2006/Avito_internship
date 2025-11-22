package postgres

import (
	"errors"
	"time"

	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
	projerrors "github.com/ArtemST2006/Avito_internship/backend/pkg/errors"
	"gorm.io/gorm"
)

type PullRequestRepo struct {
	db *gorm.DB
}

func NewPullRequestRepo(db *gorm.DB) *PullRequestRepo {
	return &PullRequestRepo{
		db: db,
	}
}

func (p *PullRequestRepo) MergePR(req schemas.PullRqMergeRequest) (schemas.PullRqResponse, error) {
	prId := req.PullRequestId
	pr := schemas.PullRequest{}

	result := p.db.Where("pull_request_id = ?", prId).First(&pr)
	if result.Error != nil {
		return schemas.PullRqResponse{}, result.Error
	}

	if pr.Status == "OPEN" {
		pr.Status = "MERGED"
		pr.MergedAt = time.Now()

		updateResult := p.db.Model(&schemas.PullRequest{}).Where("pull_request_id = ?", prId).
			Updates(map[string]interface{}{
				"status":    "MERGED",
				"merged_at": time.Now(),
			})
		if updateResult.Error != nil {
			return schemas.PullRqResponse{}, updateResult.Error
		}
	}

	response := schemas.PullRqResponse{
		Pr: pr,
	}

	return response, nil
}

func (p *PullRequestRepo) CreatePR(req schemas.CreatePullRequestRequest) (schemas.PullRqResponse, error) {
	var response schemas.PullRqResponse

	var existingPR schemas.PullRequest
	err := p.db.Where("pull_request_name = ?", req.PullRequestName).First(&existingPR).Error
	if err == nil {
		return response, projerrors.ErrAlreadyExist
	}

	var author schemas.User
	err = p.db.Where("user_id = ?", req.AuthorID).First(&author).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, projerrors.ErrNotFound
		}
		return response, err
	}

	var potentialReviewers []schemas.User
	err = p.db.Where(
		"team_name = ? AND user_id != ? AND is_active = ?",
		author.TeamName,
		req.AuthorID,
		true,
	).Limit(2).Find(&potentialReviewers).Error
	if err != nil {
		return response, err
	}

	reviewerIDs := make([]string, 0, len(potentialReviewers))
	for _, reviewer := range potentialReviewers {
		reviewerIDs = append(reviewerIDs, reviewer.UserId)
	}

	now := time.Now()
	newPR := schemas.PullRequest{
		PullRequestID:     req.PullRequestID,
		PullRequestName:   req.PullRequestName,
		AuthorID:          req.AuthorID,
		Status:            "OPEN",
		AssignedReviewers: reviewerIDs,
		CreatedAt:         now,
	}

	err = p.db.Create(&newPR).Error
	if err != nil {
		return response, err
	}

	response.Pr = newPR

	return response, nil
}

func (p *PullRequestRepo) ChangeAuthorPR(req schemas.PRChangeAuthorRequest) (schemas.PRChangeAuthorResponse, error) {
	var response schemas.PRChangeAuthorResponse

	var pr schemas.PullRequest
	err := p.db.Where("pull_request_id = ?", req.PullRequestID).First(&pr).Error
	if err != nil {
		return response, err
	}

	if pr.Status == "MERGED" {
		return response, projerrors.ErrAlreadyMerged
	}

	var oldReviewer schemas.User
	err = p.db.Where("user_id = ?", req.OldUserID).First(&oldReviewer).Error
	if err != nil {
		return response, err
	}

	var newReviewer schemas.User
	err = p.db.Where(
		"team_name = ? AND user_id != ? AND is_active = ? AND user_id NOT IN ?",
		oldReviewer.TeamName,
		req.OldUserID,
		true,
		pr.AssignedReviewers,
	).First(&newReviewer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, projerrors.ErrNoCandidate
		}
		return response, err
	}

	oldReviewerIndex := -1
	for i, reviewerID := range pr.AssignedReviewers {
		if reviewerID == req.OldUserID {
			oldReviewerIndex = i
			break
		}
	}

	if oldReviewerIndex == -1 {
		return response, errors.New("old reviewer not found in PR reviewers list")
	}

	newAssignedReviewers := make([]string, len(pr.AssignedReviewers))
	copy(newAssignedReviewers, pr.AssignedReviewers)
	newAssignedReviewers[oldReviewerIndex] = newReviewer.UserId

	pr.AssignedReviewers = newAssignedReviewers
	err = p.db.Save(&pr).Error
	if err != nil {
		return response, err
	}

	response.Pr = pr
	response.ReplacedBy = newReviewer.UserId

	return response, nil
}
