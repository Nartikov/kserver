package models

type Answer struct {
	Id        int    `json:"-" db:"id"`
	Text      string `json:"text" binding:"required"`
	IsCorrect bool   `json:"iscorrect" binding:"required"`
}
