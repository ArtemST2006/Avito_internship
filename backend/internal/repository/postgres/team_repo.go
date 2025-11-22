package postgres

import (
	"time"

	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
	projerrors "github.com/ArtemST2006/Avito_internship/backend/pkg/errors"
	"gorm.io/gorm"
)

type TeamPostgres struct {
	db *gorm.DB
}

func NewTeamRepo(db *gorm.DB) *TeamPostgres {
	return &TeamPostgres{
		db: db,
	}
}

func (t *TeamPostgres) GetTeam(teamName string) (schemas.GetTeamResponse, error) {
	response := schemas.GetTeamResponse{}

	var team schemas.Teams

	err := t.db.Where("team_name = ?", teamName).First(&team).Error
	if err != nil {
		return response, err
	}

	var users []schemas.User
	err = t.db.Where("team_name = ?", teamName).Find(&users).Error
	if err != nil {
		return response, err
	}

	response.TeamName = teamName
	response.Members = make([]schemas.Member, 0, len(users))
	for _, user := range users {
		member := schemas.Member{
			UserID:   user.UserId,
			UserName: user.UserName,
			IsActive: user.IsActive,
		}
		response.Members = append(response.Members, member)
	}

	return response, nil
}

func (t *TeamPostgres) AddTeam(req schemas.CreateTeamRequest) (schemas.CreateTeamResponse, error) {
	response := schemas.CreateTeamResponse{}

	var existingTeam []schemas.User
	err := t.db.Where("team_name = ?", req.TeamName).First(&existingTeam).Error
	if err == nil {
		return response, projerrors.ErrAlreadyExist
	}

	usersToCreate := make([]schemas.User, 0, len(req.Members))
	for _, member := range req.Members {
		user := schemas.User{
			UserId:   member.UserID,
			UserName: member.UserName,
			TeamName: req.TeamName,
			IsActive: member.IsActive,
		}
		usersToCreate = append(usersToCreate, user)
	}

	teamToCreate := schemas.Teams{
		TeamName:  req.TeamName,
		CreatedAt: time.Now(),
	}

	err = t.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&teamToCreate).Error; err != nil {
			return err
		}

		if err := tx.Create(&usersToCreate).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return response, err
	}

	response.Team.TeamName = req.TeamName
	response.Team.Members = req.Members

	return response, nil
}
