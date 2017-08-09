# todolist
该项目采用golang语言，MySql数据库，beego框架。

## 网页截图
### 登录界面
>![1](http://img.blog.csdn.net/20170809172015268?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvTW91cm5leQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

### 主页
> ![2](http://img.blog.csdn.net/20170809172045504?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvTW91cm5leQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

### 分类
> ![3](http://img.blog.csdn.net/20170809172240646?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvTW91cm5leQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

### 任务
> ![4](http://img.blog.csdn.net/20170809172121215?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvTW91cm5leQ==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

## 项目结构
```
* /conf             应用配置文件
* /controllers      所有服务端控制器
* /models           数据模型
* /routers          数据路由
* /static           静态文件
* /tests            测试文件
* /views            网页模板
* main.go           web应用程序入口
```

## 使用

```
* 将 todolist.sql 文件导入数据库，创建todolist数据库表，密码：123456
* git clone https://git.oschina.net/godev/aetchealth.git
* go run main.go
```
打开浏览器，输入 http://localhost:1234, 查看结果。