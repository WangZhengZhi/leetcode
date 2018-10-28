package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//与数据库有关的问题，表的设计有关
type User struct {
	Id   int
	Name string `orm:"unique"`//设置唯一,用户名是唯一的
	Psw  string
}

//文章结构体
type Artical struct {
	Id       int `orm:"pk;auto"` //设置为主键且自增
	Aname    string
	Atime    time.Time `orm:"auto_now_add"`
	Acount   int
	Acontent string
	Aimg     string //存一个路径就行
	//TestName string
}

func init() {
	// 设置数据库基本信息,连接数据库
	orm.RegisterDataBase("default", "mysql", "root:@163@tcp(127.0.0.1:3306)/test?charset=utf8")
	// 映射model数据
	orm.RegisterModel(new(User), new(Artical)) //create table
	// 生成表
	orm.RunSyncdb("default", false, true) //别名，强制更新，是否可见

}
