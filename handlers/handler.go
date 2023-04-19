package handlers

import (
	"github.com/gin-gonic/gin"
	"new_project/services"
)

type Hanler struct {
	service *services.Service
}

func NewHandler(services *services.Service) *Hanler {
	return &Hanler{
		service: services,
	}
}
func (h *Hanler) GetAllArticles(c *gin.Context) {
	h.service.ArticlesServier.ListAndSaveArticles()
	articles, err := services.ListArticles()
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"data":    articles,
		"message": "success",
	})
}
