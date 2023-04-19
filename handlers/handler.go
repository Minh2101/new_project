package handlers

import (
	"github.com/gin-gonic/gin"
	"new_project/services"
)

type Handler struct {
	service *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		service: services,
	}
}
func (h *Handler) GetAllArticles(c *gin.Context) {
	err := h.service.ArticlesServier.ListAndSaveArticles()
	if err != nil {
		return
	}
	articles, err := services.ListArticles()
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"data":    articles,
		"message": "success",
	})
}
