package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"goweb/apiserver_demo/demo09_X/handler"
	"goweb/apiserver_demo/demo09_X/model"
	"goweb/apiserver_demo/demo09_X/pkg/auth"
	"goweb/apiserver_demo/demo09_X/pkg/errno"
	"goweb/apiserver_demo/demo09_X/util"
)

// Create creates a new user account.
func Create(c *gin.Context) {

	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r User

	var err error
	if err = c.Bind(&r); err != nil { // 若解析 r 失败，则给客户端发送响应信息
		handler.SendResponse(c, errno.ErrBind, nil)
		log.Errorf(errno.ErrBind, "")
		return
	}

	username := c.Param("username") //url地址 ？ 前的参数
	log.Infof("URL key param desc: %s", username)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if err = r.checkParam(); err != nil { //验证用户名和密码
		handler.SendResponse(c, err, nil)
		log.Errorf(err, "get an error")
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	//selfDB := model.GetSelfDB()
	//sqlStr := "INSERT INTO dbtest VALUES(?,?)"
	model.DB.Self.LogMode(true) //开启 sql debug 模型

	r.Password, err = auth.Encrypt(r.Password)

	if err == nil {
		model.DB.Self.Create(&r)
	}

	handler.SendResponse(c, nil, rsp) //登陆成功，返回响应信息
}
