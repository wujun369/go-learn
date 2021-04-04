package UserDao

import (
	"database/sql"
	"demo04/DbTest/User"
	"fmt"
)

type UserDao struct{}

func (userDao *UserDao) AddUser(u User.User, db *sql.DB) error { //添加User
	sqlStr := "INSERT users(username,password,email) values(?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现错误：", err)
	}
	_, err1 := stmt.Exec(u.Username, u.Password, u.Email)
	if err != nil {
		fmt.Println("执行出现错误：", err1)
	}
	return nil
}

func (userDao *UserDao) GetAllUser(db *sql.DB) []User.User { //查询所有用户

	var users []User.User

	sqlStr := "SELECT * FROM users"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("预编译出现错误：", err)
	}

	for rows.Next() {
		var user User.User
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}

func (userDao *UserDao) DeleteUser(id int, db *sql.DB) { //删除用户
	sqlStr := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除成功")
	}
}

func (userDao *UserDao) UpdateUser(user User.User, db *sql.DB) {
	sqlStr := "UPDATE users SET Username = ?,Password = ?,Email = ? WHERE id = ?"
	_, err := db.Exec(sqlStr, user.Username, user.Password, user.Email, user.ID)
	if err != nil {
		fmt.Println("修改失败")
	} else {
		fmt.Println("修改成功")
	}
}
