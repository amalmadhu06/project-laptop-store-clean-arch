package middleware

import (
	"fmt"
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
	fmt.Println("user id in user auth middleware:", userID)
	c.Set("userID", userID)
	fmt.Println("user id set in auth", c.Value("userID"))
	c.Next()
}
