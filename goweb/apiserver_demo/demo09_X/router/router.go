package router

import (
	"github.com/gin-gonic/gin"
	"goweb/apiserver_demo/demo09_X/handler/sd"
	"goweb/apiserver_demo/demo09_X/handler/user"
	"goweb/apiserver_demo/demo09_X/router/middleware"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.POST("/login",user.Login)

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create) //创建用户
		u.DELETE("/:id",user.Delete) //删除用户
		u.PUT("/:id",user.Update) //更新用户
		u.GET("",user.List)
		u.GET("/:username",user.GetOne)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
