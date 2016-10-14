package controllers

import (
	"github.com/ninefive/ninefive-blog/utils"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	this.TplName = "login/login.html"
}

func (this *LoginController) DoLogin() {
	username := this.GetString("username")
	password := this.GetString("password")
	user, err := utils.CheckLogin(username, password)
}
