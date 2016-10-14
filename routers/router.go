package routers

import (
	"github.com/ninefive/ninefive-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
