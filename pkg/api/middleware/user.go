package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserAuth(c *gin.Context) {
	tokenString, err := c.Cookie("UserAuth")
	//Todo : check if user is blocked in database

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userID, err := ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("userID", userID)
	c.Next()
}
