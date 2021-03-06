package service

import (
	"github.com/nartikov/kserver/pkg/models"
	"github.com/nartikov/kserver/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Quiz interface {
	Create(userId int, quiz models.Quiz) (int, error)
}

type Question interface {
}

type Answer interface {
}

type Service struct {
	Authorization
	Quiz
	Question
	Answer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Quiz:          NewQuizService(repos.Quiz),
	}
}
