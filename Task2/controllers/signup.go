package controllers
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-contrib/sessions"
	"Task2/models"
)

func RenderSignupPage(c *gin.Context) {
    c.HTML(http.StatusOK, "signup.html", nil)
}  


func HandleSignup(c *gin.Context) {
    var form struct {
        Username string `form:"username" binding:"required"`
        Email    string `form:"email" binding:"required"`
        Password string `form:"password" binding:"required"`
    }
    if err := c.ShouldBind(&form); err != nil {
        c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Invalid form submission"})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Internal server error"})
        return
    }

    user := models.User{
        Username: form.Username,
        Email:    form.Email,
        Password: string(hashedPassword),
    }
    models.Users = append(models.Users, user)

    session := sessions.Default(c)
    session.Set("user", gin.H{
        "username": user.Username,
        "email":    user.Email,
    })
    session.Save()

    c.Redirect(http.StatusFound, "/profile")
}
