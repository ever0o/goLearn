package main

import (
	"github.com/astaxie/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

}

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:1234@/test?charset=utf8", 30)

	//注册定义的model

}
