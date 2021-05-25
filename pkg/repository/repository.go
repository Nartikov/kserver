package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nartikov/kserver/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Quiz interface {
	Create(userId int, quiz models.Quiz) (int, error)
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
		Quiz:          NewQuizPostgres(db),
	}
}
