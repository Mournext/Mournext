package routers

import (
	"todolist/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/task", &controllers.TaskController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/icon", &controllers.IconController{})
	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.CategoryController{})

}
