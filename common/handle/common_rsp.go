package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonRsp(c *gin.Context, rsp interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func RespError(c *gin.Context, code int, err error) {
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
	}
}
