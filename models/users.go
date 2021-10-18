package models

import (
	"time"
)

type User struct {
	Id int			`gorm:"column:user_id"`
	Email string
	Password string
	CreatedAt time.Time
}

type ChangeUser struct {
	Email string		`json:"email"`
	Password string		`json:"password"`
	NewPassword string	`json:"newPassword"`
}
