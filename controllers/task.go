package controllers

import (
	"todolist/models"

	"github.com/astaxie/beego"
	//"log"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Get() {

	tasks, err := models.GetAllTasks("category", false)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Tasks"] = tasks
	}

	c.TplName = "task.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	tasks, err = models.GetAllTasks(c.Input().Get("category"), true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Tasks"] = tasks

	/*   categories,err:=models.GetAllCategories()
	     if err != nil{
	      beego.Error(err)
	     }

	     c.Data["Categories"]=categories */
}

func (l *TaskController) Post() {
	//检测登录
	if !checkAccount(l.Ctx) {
		l.Redirect("/login", 302)
		return
	}

	tid := l.Input().Get("tid")
	title := l.Input().Get("title")
	//cdate := l.Input().Get("cdate")
	category := l.Input().Get("category")
	attach := l.Input().Get("attach")
	content := l.Input().Get("content")

	var err error
	if len(tid) == 0 {
		err = models.AddTask(title, category, attach, content)
	} else {
		err = models.ModifyTask(tid, title, category, attach, content)

	}

	if err != nil {
		beego.Error(err)
	}

	l.Redirect("/task", 302)
}

func (k *TaskController) Add() {

	k.Data["IsLogin"] = checkAccount(k.Ctx)
	k.TplName = "task_add.html"

}

func (t *TaskController) View() {
	t.TplName = "task_view.html"
	t.Data["IsLogin"] = checkAccount(t.Ctx)

	task, err := models.GetTask(t.Ctx.Input.Param("0"))

	if err != nil {
		beego.Error(err)
		t.Redirect("/task", 302)
		return
	}

	t.Data["Task"] = task
	t.Data["Tid"] = t.Ctx.Input.Param("0")
}

func (m *TaskController) Modify() {
	m.TplName = "task_modify.html"
	m.Data["IsLogin"] = checkAccount(m.Ctx)

	tid := m.Ctx.Input.Param("0")
	task, err := models.GetTask(tid) //获取已有内容
	if err != nil {
		beego.Error(err)
		m.Redirect("/task", 302)
		return
	}
	m.Data["Task"] = task
	m.Data["Tid"] = tid //供post操作
}

func (d *TaskController) Delete() {
	if !checkAccount(d.Ctx) {
		d.Redirect("/login", 302)
		return
	}

	err := models.DeleteTask(d.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	d.Redirect("/task", 302)
}

func (t *TaskController) Type() {
	t.TplName = "category_view.html"

	task, err := models.GetAllTasks(t.Ctx.Input.Param("0"), true)
	if err != nil {
		beego.Error(err)
		t.Redirect("/task", 302)
		return
	}

	t.Data["Task"] = task
	t.Data["IsLogin"] = checkAccount(t.Ctx)
}
