package router

import (
	"PoWTest/router/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
加载路由
*/
func Load(g *gin.Engine,mw ... gin.HandlerFunc)  *gin.Engine{

	g.Use(gin.Recovery())
	g.Use(mw...)

	//404 响应信息
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound,"页面找不到")
	})

	g.GET("/",handler.HandleGetBlockchain) //创建区块
	g.POST("/",handler.HandleWriteBlock) //查询区块

	return g
}
