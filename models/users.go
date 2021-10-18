package models

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	
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

type Token struct {
	Email string
	jwt.StandardClaims
}
