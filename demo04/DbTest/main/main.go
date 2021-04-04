package main

import (
	"database/sql"
	"demo04/DbTest/User"
	"demo04/DbTest/UserDao"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() { // = 为变量赋值；:= 创建新的变量并赋值
	var err error
	db, err = sql.Open("mysql", "root:Wujun123@tcp(localhost:3306)/dbtest")
	if err != nil {
		fmt.Println("预编译出错：", err)
	}
}

func main() {
	user := User.User{
		ID:       2,
		Username: "zhangsan",
		Password: "123456",
		Email:    "123@qq",
	}
	var userDao UserDao.UserDao

	userDao.UpdateUser(user, db)

	userDao.GetAllUser(db)
	users := userDao.GetAllUser(db)
	for _, user := range users {
		fmt.Println(user)
	}
}
