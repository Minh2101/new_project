package services

import "new_project/repository"

type Service struct {
	ArticlesServier
}
type ArticlesServier interface {
	ListAndSaveArticles() error
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ArticlesServier: NewArticlesService(repo),
	}
}
