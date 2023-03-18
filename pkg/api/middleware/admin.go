package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminAuth(c *gin.Context) {
	tokenString, err := c.Cookie("AdminAuth")

	//Todo : check if admin is blocked in database

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	adminID, err := ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("adminID", adminID)
	c.Next()
}
