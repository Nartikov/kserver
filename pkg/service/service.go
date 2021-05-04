package service

import (
	"github.com/nartikov/kserver/pkg/models"
	"github.com/nartikov/kserver/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Quiz interface {
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
	}
}
