package repository

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

func NewRepository() *Repository {
	return &Repository{}
}