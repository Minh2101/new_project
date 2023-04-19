package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"new_project/models"
	"new_project/repository"
)

type ArticlesService struct {
	ArticleRepository repository.ArticleRepoer
}

func NewArticlesService(repo *repository.Repository) *ArticlesService {
	return &ArticlesService{
		ArticleRepository: repo.ArticleRepoer,
	}
}
func ListArticles() (*models.Response, error) {
	var response models.Response
	// Tạo HTTP request
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=96a679ea1cdd4943811b159c92b58d8c")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Lấy dữ liệu từ response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *ArticlesService) ListAndSaveArticles() error {
	articles, err := ListArticles()
	if err != nil {
		return err
	}
	if err = a.ArticleRepository.SaveArticles(&articles.Articles); err != nil {
		return err
	}
	return nil
}
