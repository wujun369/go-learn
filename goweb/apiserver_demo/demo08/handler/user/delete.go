package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"goweb/apiserver_demo/demo08/handler"
	"goweb/apiserver_demo/demo08/model"
	"goweb/apiserver_demo/demo08/pkg/errno"
	"goweb/apiserver_demo/demo08/util"
	"strconv"
)

func Delete(c *gin.Context)  {
	log.Infof("User Delete function called.",lager.Data{"X-Request-Id": util.GetReqID(c)})

	id,_ := strconv.Atoi(c.Param("id"))

	user := User{}
	user.Id = uint16(id)

	model.DB.Self.LogMode(true)

	err := model.DB.Self.Delete(user).Error

	if err != nil {
		handler.SendResponse(c,errno.ErrDatabase,nil)
		return
	}

	handler.SendResponse(c,nil,nil)
}