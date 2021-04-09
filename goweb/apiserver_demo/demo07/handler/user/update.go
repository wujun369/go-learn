package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"goweb/apiserver_demo/demo07/handler"
	"goweb/apiserver_demo/demo07/model"
	"goweb/apiserver_demo/demo07/pkg/auth"
	"goweb/apiserver_demo/demo07/pkg/errno"
	"goweb/apiserver_demo/demo07/util"
	"strconv"
)

func Update(c *gin.Context)  {
	log.Infof("User Update function called.",lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r User

	if err := c.Bind(&r);err != nil{
		handler.SendResponse(c,errno.ErrBind,nil)
		return
	}

	userId, _ := strconv.Atoi(c.Param("id"))

	r.Id = uint16(userId)

	if err := r.checkParam();err != nil {
		handler.SendResponse(c,err,nil)
		return
	}

	r.Password, _ = auth.Encrypt(r.Password)

	model.DB.Self.LogMode(true)
	model.DB.Self.Save(&r)

}
