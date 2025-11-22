package repository

import "gorm.io/gorm"

//inter

type Repository struct {
	//inter
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		//inter(db)
	}
}
