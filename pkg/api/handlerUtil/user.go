package handlerUtil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("userID")
	adminID, err := strconv.Atoi(fmt.Sprintf("%v", id))
	return adminID, err
}
