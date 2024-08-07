package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"Task2/controllers"
)
func InitRouter(r *gin.Engine) {
    store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("mysession", store))

    auth := r.Group("/")
    {
        auth.GET("/login", controllers.RenderLoginPage)
        auth.POST("/login", controllers.HandleLogin)
        auth.GET("/signup", controllers.RenderSignupPage)
        auth.POST("/signup", controllers.HandleSignup)
        auth.GET("/profile", controllers.RenderProfilePage)
        auth.GET("/logout", controllers.Logout)
    }
}