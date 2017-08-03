package controllers

import (
	"todolist/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {

	// 检查是否有操作
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "category.html"

	var err error
	c.Data["Tasks"], err = models.GetAllTasks(c.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err)
	}

	//取出所有分类
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

}

func (t *CategoryController) Type() {
	t.TplName = "category_view.html"

	task, err := models.GetAllTasks(t.Ctx.Input.Param("0"), true) //Input().Get("cate")
	if err != nil {
		beego.Error(err)
		t.Redirect("/task", 302)
		return
	}

	t.Data["Task"] = task
	t.Data["IsLogin"] = checkAccount(t.Ctx)

}
