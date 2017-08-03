package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	// 判断是否为退出操作
	if c.Input().Get("exit") == "true" {
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/", 302)
		return
	}

	c.TplName = "login.html"
}

func (l *LoginController) Post() {
	//获取表单信息
	uname := l.Input().Get("username")
	pwd := l.Input().Get("password")
	autoLogin := l.Input().Get("remember") == "on"
	
	 //验证用户名及密码
	if uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPwd") {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		l.Ctx.SetCookie("username", uname, maxAge, "/")
		l.Ctx.SetCookie("password", pwd, maxAge, "/")
	}

	l.Redirect("/", 302)
	return
}


func checkAccount(ctx *context.Context) bool {

	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	
	uname := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}

	pwd := ck.Value
	return uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPwd")
}
