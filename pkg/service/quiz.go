package service

import (
	"github.com/nartikov/kserver/pkg/models"
	"github.com/nartikov/kserver/pkg/repository"
)

type QuizService struct {
	repo repository.Quiz
}

func NewQuizService(repo repository.Quiz) *QuizService {
	return &QuizService{repo: repo}
}

func (s *QuizService) Create(userId int, quiz models.Quiz) (int, error) {
	return s.repo.Create(userId, quiz)
}
