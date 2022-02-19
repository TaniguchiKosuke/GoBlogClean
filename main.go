package main

import (
	"github.com/gin-gonic/gin"

	"GoBlogClean/injector"
	article_handler "GoBlogClean/article/handler"
	auth_handler "GoBlogClean/auth/handler"
)

func main() {
	articleHandler := injector.InjectArticleHandler()
	authHandler := injector.InjectUserHandler()
	engine := gin.Default()
	article_handler.InitArticleRouting(engine, &articleHandler)
	auth_handler.InitUserRouting(engine, &authHandler)
	engine.Run(":8080")
}