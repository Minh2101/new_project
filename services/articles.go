package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"new_project/models"
)

func ListArticles(c *gin.Context) {
	var response models.Response
	// Tạo HTTP request
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=96a679ea1cdd4943811b159c92b58d8c")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Lấy dữ liệu từ response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(body, &response); err != nil {
		fmt.Println(err)
		return
	}

	// Hiển thị dữ liệu lấy được
	c.AbortWithStatusJSON(200, gin.H{
		"code": 200,
		"data": response,
	})
}
