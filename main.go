package main

import (
	"github.com/gin-gonic/gin"

	"GoBlogClean/injector"
	"GoBlogClean/article/handler"
)

func main() {
	articleHandler := injector.InjectArticleHandler()
	engine := gin.Default()
	handler.InitArticleRouting(engine, &articleHandler)
	engine.Run(":8080")
}