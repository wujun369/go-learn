package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"goweb/apiserver_demo/demo07/handler"
	"goweb/apiserver_demo/demo07/model"
	"goweb/apiserver_demo/demo07/pkg/errno"
	"goweb/apiserver_demo/demo07/util"
)

func GetOne(c *gin.Context) {

	log.Infof("User GetOne function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	username := c.Param("username")

	if username == "" { //如果 username 为空，调用 List 方法
		List(c)
	}

	user := &User{}

	model.DB.Self.LogMode(true)
	model.DB.Self.Where("username = ?", username).First(&user)

	if user == nil { //查询失败，返回响应信息
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, user)

}
