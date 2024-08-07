package article
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task_1/models"

)
func UpdateArticle(c *gin.Context) {
    id := c.Param("id")
    var updatedArticle models.Article
    if err := c.ShouldBindJSON(&updatedArticle); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, ar:= range models.Articles {
        if ar.ID == id {
            models.Articles[i] = updatedArticle
            c.JSON(http.StatusOK, updatedArticle)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}
