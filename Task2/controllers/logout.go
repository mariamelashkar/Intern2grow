package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()

    if err := session.Save(); err != nil {
        c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to save session"})
        return
    }

    c.Redirect(http.StatusFound, "/login")
}
