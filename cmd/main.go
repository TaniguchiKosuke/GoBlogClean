package main

import (
	"github.com/gin-gonic/gin"

	article_handler "GoBlogClean/article/handler"
	auth_handler "GoBlogClean/auth/handler"
	"GoBlogClean/injector"
)

func main() {
	articleHandler := injector.InjectArticleHandler()
	authHandler := injector.InjectUserHandler()
	engine := gin.Default()
	article_handler.InitArticleRouting(engine, &articleHandler)
	auth_handler.InitUserRouting(engine, &authHandler)
	engine.Run(":8080")
}
