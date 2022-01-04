package helpers

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Roles string `json:"roles"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
