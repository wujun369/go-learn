package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"goweb/apiserver_demo/demo06/handler"
	"goweb/apiserver_demo/demo06/pkg/errno"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var r CreateRequest

	var err error
	if err := c.Bind(&r); err != nil {// 若解析 r 失败，则给客户端发送响应信息
		handler.SendResponse(c, errno.ErrBind, nil)
		log.Errorf(errno.ErrBind,"")
		return
	}

	username := c.Param("username")//url地址 ？ 前的参数
	log.Infof("URL key param desc: %s", username)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s",desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx"))
		handler.SendResponse(c, err, nil)//向客户端发送错误响应

		log.Errorf(err, "Get an error")//打印日志：用户在XX数据库中找不到
		return
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
		return
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
		handler.SendResponse(c,err,nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c,nil,rsp)//登陆成功，返回响应信息
}
