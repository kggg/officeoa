package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:rootpass@tcp(10.68.2.30:3306)/cnnoc?charset=utf8")
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.Debug = true
}
