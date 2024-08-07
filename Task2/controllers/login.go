package controllers
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"Task2/models"
	"github.com/gin-contrib/sessions"
	"log"

)
func RenderLoginPage(c *gin.Context) {
    c.HTML(http.StatusOK, "login.html", nil)
}

func HandleLogin(c *gin.Context) {
    var form struct {
        Username string `form:"username" binding:"required"`
        Password string `form:"password" binding:"required"`
    }
    if err := c.ShouldBind(&form); err != nil {
        c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid form submission"})
        return
    }

    for _, user := range models.Users {
        if user.Username == form.Username {
            if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
                c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
                return
            }

            session := sessions.Default(c)
            session.Set("user", gin.H{
                "username": user.Username,
                "email":    user.Email,
            })

            log.Println("Session data before saving:", session.Get("user")) // Debug statement 

            if err := session.Save(); err != nil {
                log.Println("Error saving session:", err) // Debug statement 
                c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Failed to save session"})
                return
            }

            log.Println("Session set for user:", user.Username) // Debug statement

            c.Redirect(http.StatusFound, "/profile")
            return
        }
    }

    c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
}