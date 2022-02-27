package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"GoBlogClean/pkg/constant"
)

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

func CreateJWTToken(userID string, username string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"sub":  userID,
		"name": username,
		"iat":  time.Now(),
		"exp":  time.Now().Add(time.Minute * constant.TokenExpirationMinute).Unix(),
	}

	jwtSecretKey := jwtSecretKey

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetUserIDFromJWT(tokenStr string) (string, error) {
	token, err := VerifyToken(tokenStr)
	if err != nil {
		return "", err
	}

	jwtClaims := token.Claims.(jwt.MapClaims)
	return jwtClaims["sub"].(string), nil
}
