package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nartikov/kserver/pkg/models"
)

type QuizPostgres struct {
	db *sqlx.DB
}

func NewQuizPostgres(db *sqlx.DB) *QuizPostgres {
	return &QuizPostgres{db: db}
}

func (r *QuizPostgres) Create(userId int, quiz models.Quiz) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", quizzesTable)
	row := tx.QueryRow(createQuery, quiz.Name, quiz.Picture)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersQuizQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersQuizzes)
	_, err = tx.Exec(createUsersQuizQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
