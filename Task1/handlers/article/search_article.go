package article
 import (
	"task_1/models"
	"errors"
 )

func SearchForArticle(id string) (*models.Article, error) { 

	for index, value := range models.Articles {
		if value.ID == id {
			return &models.Articles[index], nil 
		}
	}
	return nil, errors.New("article not found")
}