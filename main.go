package main

import (
	_ "todolist/routers"
	 "todolist/models" 

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug=true
	orm.RunSyncdb("default", false, true)

	beego.Run()
}

