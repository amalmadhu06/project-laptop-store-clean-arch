package handlerUtil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAdminIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("adminID")
	adminID, err := strconv.Atoi(fmt.Sprintf("%v", id))
	return adminID, err
}
