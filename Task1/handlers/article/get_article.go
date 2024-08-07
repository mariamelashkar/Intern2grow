package article

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task_1/models"

)
func GetArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Articles)

}

func GetArticle(c *gin.Context) {
	ID := c.Param("id") 
	article, err := SearchForArticle(ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The article is Not found"}) 
		return
	}
	c.IndentedJSON(http.StatusOK, article)
}