package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"goweb/apiserver_demo/demo09_X/handler"
	"goweb/apiserver_demo/demo09_X/model"
	"goweb/apiserver_demo/demo09_X/pkg/constvar"
	"goweb/apiserver_demo/demo09_X/pkg/errno"
	"goweb/apiserver_demo/demo09_X/util"
	"sync"
)

func List(c *gin.Context) {

	log.Infof("User List function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r ListRequest

	if err := c.BindJSON(&r); err != nil {//判定绑定JSON数据是否成功
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	users, count, err := ListUserAndSync(r.Username, r.Offset,r.Limit)//并发查询
	if err != nil{
		handler.SendResponse(c,err,nil)
	}
	handler.SendResponse(c,nil,ListResponse{//响应客户端
		TotalCount: count,
		UserList: users,
	})
}

func ListUserAndSync(username string, offset int, limit int) ([]*User, uint64, error) {

	infos := make([]*User,0)
	users, count, err := ListUser(username, offset, limit)//从数据库查询所有用户

	if err != nil{
		return nil,count,err
	}

	ids := []uint64{}
	for _,user := range users{
		ids = append(ids,uint64(user.Id))
	}

	wg := sync.WaitGroup{}

	userList := UserList{
		Lock: new(sync.Mutex),
		IdMap: make(map[uint64]*User,len(users)),
	}

	errChan := make(chan error,1)
	finished := make(chan bool,1)

	/*
		以下为并发查询
	 */

	for _,u := range users{
		wg.Add(1)

		go func(u *User) {
			defer wg.Done()

			userList.Lock.Lock()
			userList.Lock.Unlock()
			userList.IdMap[uint64(u.Id)] = &User{
				Id: u.Id,
				Username: u.Username,
				Password: u.Password,
			}

		}(u)

	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <- finished:
	case err := <- errChan:
		return nil,count,err
		}

	for _,id := range ids{
		infos = append(infos,userList.IdMap[id])
	}

return infos,count,nil

}

func ListUser(username string, offset int, limit int) ([]*User, uint64, error) {

	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*User, 0)
	var count uint64

	//where := fmt.Sprintf("username like %s", username)
	if err := model.DB.Self.Model(&User{}).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := model.DB.Self.Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users,count,err
	}

	return users,count,nil

}
