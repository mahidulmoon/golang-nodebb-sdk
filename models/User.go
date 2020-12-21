package models

type User struct {
	Email string	`json:"email" binding:"required"`
	Username string	`json:"username" binding:"required"`
}
