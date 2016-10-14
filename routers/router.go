package routers

import (
	"github.com/astaxie/beego"
	"github.com/ninefive/ninefive-blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/article/:ident", &controllers.MainController{}, "get:Read")
	beego.Router("/catalog/:ident", &controllers.MainController{}, "get:ListByCatalog")

	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/admin", &controllers.AdminController{}, "get:Default")
	beego.Router("/admin/catalog/add", &controllers.CatalogController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/catalog/edit", &controllers.CatalogController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/catalog/del", &controllers.CatalogController{}, "get:Del")
	beego.Router("/admin/article/add", &controllers.ArticleController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/article/edit", &controllers.ArticleController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/article/del", &controllers.ArticleController{}, "get:Del")
	beego.Router("/admin/article/draft", &controllers.ArticleController{}, "get:Draft")
}
