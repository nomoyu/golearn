package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtCustClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id int, name string) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * 60 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	return token.SignedString([]byte("nomoyu"))

}

func ParseToken(t string) (JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(t, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("nomoyu"), nil
	})
	if err == nil && !token.Valid {
		err = errors.New("Invalid Token!")
	}
	return iJwtCustClaims, err
}

func IsTokenValid(t string) bool {
	_, err := ParseToken(t)
	if err != nil {
		return false
	}
	return true
}
