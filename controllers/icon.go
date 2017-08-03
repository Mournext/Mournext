package controllers

import (
	"github.com/astaxie/beego"
)

type IconController struct {
	beego.Controller
}

func (c *IconController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "icon.html"
}
