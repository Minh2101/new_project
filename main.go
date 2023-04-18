package main

import (
	"github.com/gin-gonic/gin"
	"new_project/services"
)

func main() {
	router := gin.Default()
	router.GET("/articles", services.ListArticles)
	router.Run(":8080")
}
