package controllers

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)
func RenderProfilePage(c *gin.Context) {
    session := sessions.Default(c)
    user := session.Get("user")
    if user == nil {
        c.Redirect(http.StatusFound, "/login")
        return
    }
    c.HTML(http.StatusOK, "profile.html", user)
}