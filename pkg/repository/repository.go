package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nartikov/kserver/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Quiz interface {
}

type Question interface {
}

type Answer interface {
}

type Repository struct {
	Authorization
	Quiz
	Question
	Answer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
