package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

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
	return &Repository{}
}