package article
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task_1/models"

)

func DeleteArticle(c *gin.Context) {
    id := c.Param("id")
    for i, a := range models.Articles {
        if a.ID == id {
            models.Articles = append(models.Articles[:i], models.Articles[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}