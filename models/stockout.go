package models

import (
	"github.com/astaxie/beego/orm"
)

type Stockout struct {
	Id         int
	Series     string
	Model      string
	Serialno   int
	Status     string
	Boxid      string
	Owner      string
	Created_at string
	Remark     string
}

func init() {
	orm.RegisterModel(new(Stockout))
}

func FindAllStockout() ([]Stockout, error) {
	var stock []Stockout
	o := orm.NewOrm()
	_, err := o.QueryTable("stockout").All(&stock)
	return stock, err
}

func FindStockoutBySerialno(num int) ([]Stockout, error) {
	var stock []Stockout
	o := orm.NewOrm()
	_, err := o.QueryTable("stockout").Filter("serialno", num).All(&stock)
	return stock, err
}

func AddStockout(series, model string, serialno int, status, boxid, owner, remark string, o orm.Ormer) (int64, error) {
	sql := "insert into stockout (series,model,serialno,status, boxid,owner, remark) values(?,?,?,?,?,?,?)"
	res, err := o.Raw(sql, series, model, serialno, status, boxid, owner, remark).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}
