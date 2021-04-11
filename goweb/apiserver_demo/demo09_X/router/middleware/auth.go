package middleware

import (
	"github.com/gin-gonic/gin"
	"goweb/apiserver_demo/demo09_X/handler"
	"goweb/apiserver_demo/demo09_X/pkg/errno"
	"goweb/apiserver_demo/demo09_X/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
