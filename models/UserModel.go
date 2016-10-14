package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	Username string
	Password string
}

func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "Username")
	return user
}
