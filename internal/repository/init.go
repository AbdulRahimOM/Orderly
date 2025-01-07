package repo

import (
	repositoryinterface "orderly/internal/repository/interface"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositoryinterface.Repository {
	return &Repo{db: db}
}
