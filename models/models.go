package models

import (
	"strconv"
	"time"
	//"string"

	"github.com/astaxie/beego/orm"
	//“_”指只执行驱动初始化
	_ "github.com/go-sql-driver/mysql"

)

var (	 	
	// 设置数据库名称
	_MYSQL_DRIVER = "mysql"
 ) 


// 分类
type Category struct {
	Id              int64
	Title           string
	TaskCount      int64
}

// 任务
type Task struct {
	Id             		int64
	Title          		string
	Cate 	          	string
	Content        		string `orm:"size(5000)"`	
	Attachment      	string	
	Created_date   		time.Time `orm:"index"`
	Last_modified_at  	time.Time `orm:"index"`
	Finish_date			time.Time `orm:"index"`
	
}

func RegisterDB() {
	
	// 注册模型
	orm.RegisterModel(new(Category),new(Task))
	
	// 注册驱动
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)
	
	// 注册默认数据库
	orm.RegisterDataBase("default",_MYSQL_DRIVER, "root:123456@/todolist?charset=utf8")
}

func AddTask(title,category,attch,content string) error {//cdate,
	//category = "$"+string.Join()
	o := orm.NewOrm()
	
	task := &Task{
		Title:			title,
		Cate:	  		category,
		Attachment:		attch,
		Content:		content,
		Created_date:   time.Now(),
		Last_modified_at: time.Now(),
	}

	_,err:=o.Insert(task)
	if err != nil {
		return err
	}

	// 更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		// 如果不存在就直接忽略，只当分类存在时进行更新
		cate.TaskCount++
		_, err = o.Update(cate)
	}

	return err
	
}

func GetAllTasks(category string, isDesc bool) (tasks []*Task, err error) {
	o := orm.NewOrm()

	tasks = make([]*Task, 0)

	qs := o.QueryTable("task")

	if isDesc {
		if len(category) > 0 {
			qs = qs.Filter("cate",category)
		}
		_, err = qs.OrderBy("-created_date").All(&tasks)
	} else {
		_, err = qs.All(&tasks)
	}
	
	return tasks, err
}


func GetTask(tid string) (*Task, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	task := new(Task)

	qs := o.QueryTable("task")
	err = qs.Filter("id", tidNum).One(task)
	if err != nil {
		return nil, err
	}

/*	task.Views++
	_, err = o.Update(task)*/
	return task, nil   
}

func ModifyTask(tid,title,category,attach,content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)//字符串转换数字
	if err != nil {
		return err
	}

	var oldCate string
	o:=orm.NewOrm()

	task := &Task {Id : tidNum}
	if o.Read(task) == nil {
		oldCate			=	task.Cate
		task.Title 		=	title
		task.Cate		=	category		
		task.Content 	= 	content
		task.Attachment =	attach
		task.Last_modified_at	=	time.Now()
		_,err = o.Update(task)
		if err !=nil{
			return err
		}		
	}

	//更新分类统计
	if len(oldCate) > 0{
		cate := new(Category)
		qs  :=	o.QueryTable("category")
		err  = qs.Filter("title",oldCate).One(cate)
		if err == nil {
			cate.TaskCount--
			_, err =o.Update(cate)
		}		
	}

	cate := new (Category)
	qs := o.QueryTable("category")
	err =qs.Filter("title",category).One(cate)
	if err == nil {
		cate.TaskCount++
		_,err = o.Update(cate)
	}

	return nil
}

func DeleteTask(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	var oldCate string
	o := orm.NewOrm()

	task := &Task{Id: tidNum}
	if o.Read(task) == nil {
		oldCate = task.Cate
		_, err = o.Delete(task)
		if err != nil {
			return err
		}
	}
	//更新分类统计
	if len(oldCate) > 0{
		cate := new(Category)
		qs  :=	o.QueryTable("category")
		err  = qs.Filter("title",oldCate).One(cate)
		if err == nil {
			cate.TaskCount--
			_, err =o.Update(cate)
		}		
	}

	return err
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title:name}

	// 查询数据
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id string) error {
	cid,err:=strconv.ParseInt(id,10,64)
		if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
} 