package repository

import (
	"gorm.io/gorm"
)

type Repo struct {
	gorm *gorm.DB
}

type RepoInterface interface {
	BookRepo
}

func NewRepo(gorm *gorm.DB) *Repo {
	return &Repo{gorm: gorm}
}
