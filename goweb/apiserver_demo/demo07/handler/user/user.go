package user

import (
	"goweb/apiserver_demo/demo07/pkg/errno"
	"sync"
)

type User struct {
	Id uint16
	Username string		`json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string		`json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

type CreateResponse struct {
	Username string		`json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

func (r *User) checkParam() error{ //验证用户名和密码是否有效
	if r.Username == "" {
		return errno.New(errno.ErrValidation,nil).Add("username is empty")
	}
	if r.Password == ""{
		return errno.New(errno.ErrValidation,nil).Add("password is empty")
	}
	return nil
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*User
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*User `json:"userList"`
}