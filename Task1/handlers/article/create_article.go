package article

import (
	"net/http"
	"task_1/models"
	"github.com/gin-gonic/gin"
)
func CreateArticle(c *gin.Context) {
	var newArticle models.Article
	if err := c.BindJSON(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.Articles = append(models.Articles, newArticle)
	c.JSON(http.StatusCreated, newArticle)
}
