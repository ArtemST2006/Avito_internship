// Package schemas содержит структуры данных для работы с базой и JSON.
package schemas

import "time"

// User представляет сущность пользователя в системе
type User struct {
	UserId   string `gorm:"column:user_id;primaryKey" json:"user_id" binding:"required"`
	UserName string `gorm:"column:username" json:"username" binding:"required"`
	TeamName string `gorm:"column:team_name" json:"team_name" binding:"required"`
	IsActive bool   `gorm:"column:is_active" json:"is_active" binding:"required"`
}

// TeamMember представляет участника команды
type TeamMember struct {
	UserId   string `gorm:"column:user_id" json:"user_id" binding:"required"`
	UserName string `gorm:"column:username" json:"username" binding:"required"`
	IsActive bool   `gorm:"column:is_active" json:"is_active" binding:"required"`
}

// Teams представляет команду в системе
type Teams struct {
	TeamName  string    `gorm:"column:team_name;primaryKey" json:"team_name" binding:"required"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt,omitempty"`
}

// PullRequest представляет pull request в системе
type PullRequest struct {
	PullRequestID     string    `gorm:"column:pull_request_id;primaryKey" json:"pull_request_id" binding:"required"`
	PullRequestName   string    `gorm:"column:pull_request_name" json:"pull_request_name" binding:"required"`
	AuthorID          string    `gorm:"column:author_id" json:"author_id" binding:"required"`
	Status            string    `gorm:"column:status" json:"status" binding:"required"`
	AssignedReviewers []string  `gorm:"column:assigned_reviewers;type:jsonb;serializer:json" json:"assigned_reviewers"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"createdAt,omitempty"`
	MergedAt          time.Time `gorm:"column:merged_at" json:"mergedAt,omitempty"`
}

// PullRequestShort представляет сокращенную информацию о pull request
// type PullRequestShort struct {
// 	PullRequestID   string `json:"pull_request_id"`
// 	PullRequestName string `json:"pull_request_name"`
// 	AuthorID        string `json:"author_id"`
// 	Status          string `json:"status"` // "OPEN" или "MERGED"
// }
