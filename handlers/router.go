package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	h.SetTimeFeedData()
	router := gin.New()
	api := router.Group("/api/")
	api.GET("/articles", h.GetAllArticles)
	return router
}
