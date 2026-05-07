package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CommonRsp(c *gin.Context, rsp interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func Resp(c *gin.Context, code int, msg string, data interface{}) {
	rsp := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, rsp)
}

func RespError(c *gin.Context, code int, err error) {
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
	}
}
