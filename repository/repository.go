package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"new_project/models"
)

type Repository struct {
	ArticleRepoer
}
type ArticleRepoer interface {
	SaveArticles(articles *[]models.Article) error
}

func CollectionName() string {
	return "articles"
}
func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		ArticleRepoer: NewArticleRepo(db, CollectionName()),
	}
}
