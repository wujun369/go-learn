package handler

import (
	"github.com/gin-gonic/gin"
	"goweb/apiserver_demo/demo07/pkg/errno"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SendResponse(c *gin.Context,err error,data interface{})  {//向客户端发送响应信息
	code, message := errno.DecodeErr(err)

	r := Response{
		Code: code,
		Message: message,
		Data: data,
	}

	c.JSON(http.StatusOK,r)
}