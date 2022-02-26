package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"GoBlogClean/injector"
	auth_handler "GoBlogClean/pkg/auth/handler"
	article_handler "GoBlogClean/pkg/blog/handler"
)

func main() {
	articleHandler := injector.InjectArticleHandler()
	authHandler := injector.InjectUserHandler()
	engine := gin.Default()
	article_handler.InitArticleRouting(engine, &articleHandler)
	auth_handler.InitUserRouting(engine, &authHandler)
	if err := engine.Run(":8080"); err != nil {
		log.Println(err)
	}
}
