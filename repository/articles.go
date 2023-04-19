package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"new_project/database"
	"new_project/models"
	"sync"
)

type ArticleRepo struct {
	db         *mongo.Database
	collection string
}

func NewArticleRepo(db *mongo.Database, collections string) ArticleRepoer {
	return &ArticleRepo{
		db:         db,
		collection: collections,
	}
}
func (a *ArticleRepo) SaveArticles(articles *[]models.Article) error {
	var (
		collection  = a.db.Collection(a.collection)
		ctx, cancel = context.WithTimeout(context.Background(), database.CTimeOut)
	)
	defer cancel()
	// Lưu dữ liệu vào cơ sở dữ liệu
	var wg sync.WaitGroup
	for _, article := range *articles {
		wg.Add(1)
		go func(article models.Article) {
			defer wg.Done()
			//filter
			filter := bson.M{
				"url": article.URL,
			}
			//check exist articles
			if err := collection.FindOne(context.Background(), filter).Decode(&article); err == nil {
				return
			}
			if _, err := collection.InsertOne(ctx, article); err != nil {
				return
			}
		}(article)
	}
	wg.Wait()
	return nil
}
