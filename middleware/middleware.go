package middleware

import (
	"BelajarGolang5/models"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthValid(c *gin.Context) {
	tokenString, err := c.Cookie("token")

	// run if no Cookie
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "token nil"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not valid")
		}
		return []byte(models.SECRET), nil
	})

	if err == nil && token.Valid {
		c.Next()
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "token is expiry"})
		c.Abort()
	}
}
