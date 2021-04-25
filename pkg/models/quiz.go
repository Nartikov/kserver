package models

type Quiz struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Picture string `json:"picture" binding:"required"`
}