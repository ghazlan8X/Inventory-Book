package auth

import (
	"BelajarGolang5/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type authHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) authHandler {
	return authHandler{DB: db}
}

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/login")
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"content": "",
	})
}

func (h *authHandler) LoginPostHandler(c *gin.Context) {
	var input models.Login
	var user models.User

	if err := c.Bind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "login", gin.H{"content": "input not valid"})
		return
	}

	if err := h.DB.Where("username =?", input.Username).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login", gin.H{"content": "username not found"})
		return
	}

	if user.Password != input.Password {
		c.HTML(http.StatusUnauthorized, "login", gin.H{"content": "Password invalid"})
		return
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		Issuer:    "book-inventory",
		IssuedAt:  time.Now().Unix(),
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := sign.SignedString([]byte(models.SECRET))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login", gin.H{
			"contet": "internal server error",
		})
		c.Abort()
	}

	c.SetCookie("token", token, 3600*2, "/", "localhost", false, true)

	c.Redirect(http.StatusSeeOther, "/books")

}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}
