package authenticator

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var claims *jwt.MapClaims

// SetToken to simplify use of JWT auth
func SetToken(username string, appName string) *jwt.Token {

	expTime := time.Now().Add(time.Hour * time.Duration(48)).Unix() // expired time of JWT token
	issuedAt := time.Now().Unix()                                   // issued at of JWT token

	claims = &jwt.MapClaims{
		"user": username,
		"exp":  expTime,
		"iat":  issuedAt,
		"iss":  appName,
	}

	ret := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return ret
}
