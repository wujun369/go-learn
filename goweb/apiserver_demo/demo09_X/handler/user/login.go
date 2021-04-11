package user

import (
	"github.com/gin-gonic/gin"
	"goweb/apiserver_demo/demo09_X/handler"
	"goweb/apiserver_demo/demo09_X/model"
	"goweb/apiserver_demo/demo09_X/pkg/auth"
	"goweb/apiserver_demo/demo09_X/pkg/errno"
	"goweb/apiserver_demo/demo09_X/pkg/token"
)

func Login(c *gin.Context)  {

	var user User
	var database_user User

	if err := c.BindJSON(&user);err != nil{
		handler.SendResponse(c,errno.ErrBind,nil)
		return
	}

	if err := model.DB.Self.Where("username = ?", user.Username).First(&database_user).Error;err != nil{
		handler.SendResponse(c,errno.ErrUserNotFound,nil)
		return
	}

	if err := auth.Compare(database_user.Password,user.Password);err != nil{
		handler.SendResponse(c,errno.ErrPasswordIncorrect,nil)
		return
	}

	sign, err := token.Sign(c, token.Context{
		ID:       uint64(database_user.Id),
		Username: database_user.Username,
	}, "")

	if err != nil{
		handler.SendResponse(c,errno.ErrToken,nil)
	}

	handler.SendResponse(c,nil,Token{Token: sign})

}
