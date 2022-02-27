package main

import (
	"log"

	"github.com/gin-gonic/gin"

	auth_handler "GoBlogClean/pkg/auth/handler"
	article_handler "GoBlogClean/pkg/blog/handler"
	"GoBlogClean/pkg/di"
)

func main() {
	articleHandler := di.InjectArticleHandler()
	authHandler := di.InjectUserHandler()
	engine := gin.Default()
	article_handler.InitArticleRouting(engine, &articleHandler)
	auth_handler.InitUserRouting(engine, &authHandler)
	if err := engine.Run(":8080"); err != nil {
		log.Println(err)
	}
}
