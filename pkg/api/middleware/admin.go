package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminAuth(c *gin.Context) {
	tokenString, err := c.Cookie("AdminAuth")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err := ValidateToken(tokenString); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
