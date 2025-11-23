package postgres

import (
	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
	"gorm.io/gorm"
)

type StatisticPostgres struct {
	db *gorm.DB
}

func NewStatisticPostgres(db *gorm.DB) *StatisticPostgres {
	return &StatisticPostgres{
		db: db,
	}
}

func (s *StatisticPostgres) Statistic() (schemas.StatisticResponse, error) {
	var resp schemas.StatisticResponse
	var openCount, mergedCount int64

	if err := s.db.Model(&schemas.PullRequest{}).Where("status = ?", "OPEN").Count(&openCount).Error; err != nil {
		return resp, err
	}
	if err := s.db.Model(&schemas.PullRequest{}).Where("status = ?", "MERGED").Count(&mergedCount).Error; err != nil {
		return resp, err
	}

	resp.NumberOpen = int(openCount)
	resp.NumberMerged = int(mergedCount)

	type UserPRStat struct {
		UserID     string `gorm:"column:user_id"`
		NumberOfPR int    `gorm:"column:number_of_pr"`
	}

	var stats []UserPRStat
	if err := s.db.
		Table("users").
		Select("users.user_id, COUNT(pull_requests.pull_request_id) AS number_of_pr").
		Joins("LEFT JOIN pull_requests ON users.user_id = pull_requests.author_id").
		Where("users.is_active = ?", true).
		Group("users.user_id").
		Scan(&stats).Error; err != nil {
		return resp, err
	}

	resp.PRUser = make([]schemas.StatisticForUser, len(stats))
	for i, stat := range stats {
		resp.PRUser[i] = schemas.StatisticForUser{
			UserID:     stat.UserID,
			NumberOfPR: stat.NumberOfPR,
		}
	}

	return resp, nil
}
