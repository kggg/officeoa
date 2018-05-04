package models

import (
	"github.com/astaxie/beego/orm"
)

type Stock struct {
	Id       int    `json:"id"`
	Series   string `json:"series"; valid:"Required;MaxSize(15)`
	Model    string `json:"model"; valid:"Required;Match(/-/);MaxSize(20)"`
	Serialno int    `json:"serialno"; valid:"Required;Range(1000, 99999999);MaxSize(20)"`
	Status   string `json:"status"`
}

func init() {
	orm.RegisterModel(new(Stock))
}

type Stocks struct {
	Stock
	Remark string `valid:"MaxSize(300)"`
	Date   string `valid:"MaxSize(20)"`
}

func FindAllStock() ([]Stock, error) {
	var stock []Stock
	o := orm.NewOrm()
	_, err := o.QueryTable("stock").All(&stock)
	return stock, err
}

func FindStockBySerialno(num int) (Stock, error) {
	var stock Stock
	o := orm.NewOrm()
	err := o.QueryTable("stock").Filter("serialno", num).One(&stock)
	return stock, err
}

func SerialnoExistCheck(num int) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("stock").Filter("serialno", num).Exist()
	return exist
}

func StockStatusChange(serialno int, status string, o orm.Ormer) (int64, error) {
	num, err := o.QueryTable("stock").Filter("serialno", serialno).Update(orm.Params{
		"status": status,
	})
	return num, err
}

func PaginatorStock(page, size int, name string) ([]Stock, int64, error) {
	o := orm.NewOrm()
	var stock []Stock
	if name == "" {
		qs := o.QueryTable("stock")
		count, _ := qs.Count()

		_, err := qs.Limit(size).Offset(size * (page - 1)).All(&stock)
		return stock, count, err
	} else {
		cond := orm.NewCondition()
		cond1 := cond.And("series__icontains", name).Or("model__icontains", name).Or("serialno__icontains", name).Or("status__contains", name)
		qs := o.QueryTable("stock")
		qs = qs.SetCond(cond1)
		count, _ := qs.Count()
		qs = qs.Limit(size).Offset(size * (page - 1))
		_, err := qs.All(&stock)
		return stock, count, err
	}
}

func AddStock(stock *Stock, o orm.Ormer) error {
	stock.Status = "空闲机"
	_, err := o.Insert(stock)
	return err
}

func EditStock(s *Stock, id int) (int64, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("stock").Filter("id", s.Id).Update(orm.Params{
		"serialno": s.Serialno,
		"series":   s.Series,
		"model":    s.Model,
	})
	return num, err

}

func StockAdd(stocks *Stocks) error {
	o := orm.NewOrm()
	err := o.Begin()
	AddStockin(stocks.Id, stocks.Date, stocks.Remark, o)
	err = AddStock(&stocks.Stock, o)
	if err != nil {
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}

func StockDelete(id int) error {
	o := orm.NewOrm()
	err := o.Begin()
	DeleteStock(id, o)
	_, err = DeleteStockin(id, o)
	if err != nil {
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err

}

func DeleteStock(id int, o orm.Ormer) (int64, error) {
	if num, err := o.Delete(&Stock{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}
