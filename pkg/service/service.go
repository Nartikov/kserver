package service

import "github.com/nartikov/kserver/pkg/repository"

type Authorization interface {

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
	return &Service{}
}