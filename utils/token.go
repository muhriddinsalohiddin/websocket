package utils

import (
	"log"
	"time"
	"app/models"
	jwt "github.com/dgrijalva/jwt-go"
)


var jwtKey = []byte("hello_world")


func CreateToken (u models.User) (string){


	expirationTime := time.Now().Add(3 * time.Minute)

	tokenStruct := models.Token {
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

func CheckToken (token string)bool {

	tkn, err := jwt.ParseWithClaims(token, &models.Token{}, func(token *jwt.Token) (interface{}, error) {
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

