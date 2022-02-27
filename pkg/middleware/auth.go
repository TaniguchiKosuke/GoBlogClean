package middleware

import (
	"GoBlogClean/pkg/util"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTをAuthorization Headerから受け取り認証する
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Authorization is empty")
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := util.VerifyToken(tokenStr)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}
	}
}
