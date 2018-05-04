package models

import (
	"github.com/astaxie/beego/orm"
)

type Stockin struct {
	Id       int
	Stock_id int
	Indate   string
	Remark   string
}

func init() {
	orm.RegisterModel(new(Stockin))
}

func AddStockin(id int, indate, remark string, o orm.Ormer) (int64, error) {
	sql := "insert into stockin (stock_id, indate, remark) values(?,?,?)"
	res, err := o.Raw(sql, id, indate, remark).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}

func DeleteStockin(id int, o orm.Ormer) (int64, error) {
	if num, err := o.Delete(&Stockin{Stock_id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}
