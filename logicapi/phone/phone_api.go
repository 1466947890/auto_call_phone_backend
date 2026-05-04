package phone

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// c 传递的是指针
func GetPhoneList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Method": "Get"})
}
