package routers

import (
	"task_1/handlers/article"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine){

    r.GET("/article/:id", article.GetArticle)
	r.GET("/articles", article.GetArticles)
    r.POST("/article", article.CreateArticle)
    r.PUT("/article/:id", article.UpdateArticle)
    r.DELETE("/article/:id", article.DeleteArticle)

}