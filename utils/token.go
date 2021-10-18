package utils

import (
	"log"
	"time"
	"app/models"
	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	Email string
	jwt.StandardClaims
}

var jwtKey = []byte("hello_world")

func CheckToken (token string)bool {

	tkn, err := jwt.ParseWithClaims(token, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false
	}
	if !tkn.Valid {
		return false
	}

	return true
}


func CreateToken (u models.User) (string){


	expirationTime := time.Now().Add(30 * time.Minute)

	tokenStruct := Token {
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenStruct)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}


