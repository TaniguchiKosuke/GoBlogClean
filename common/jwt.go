package common

import (
	"GoBlogClean/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWTToken(user *models.User) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user_id": user.ID,
		"username": user.Username,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}