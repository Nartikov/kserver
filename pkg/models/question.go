package models

type Question struct {
	Id       int    `json:"-" db:"id"`
	Text     string `json:"text" binding:"required"`
	Picture string `json:"picture" binding:"required"`
	TimeLimit int `json:"timelimit" binding:"required"`
}