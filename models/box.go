package models

import (
	"github.com/astaxie/beego/orm"
)

type Box struct {
	Id          int    `json:"id"`
	Owner_id    int    `json:"owner_id"; valid:"Required;Range(0,100000)"`
	Installdate string `json:"installdate";`
	Boxid       string `json:"boxid"; valid:"Required; MaxSize(50)"`
	Serialno    int    `json:"serialno"; valid:"Required;Range(1000,90000000) "`
	Service     string `json:"service"; valid:"Required;MaxSize(20)"`
	Zoomowner   string `json:"zoomowner"`
	Hqowner     string `json:"hqowner"`
	Usernums    int    `json:"usernums"`
	State       string `json:"state"`
	Remark      string `json:"remark"`
}

func init() {
	orm.RegisterModel(new(Box))
}

func FindAllBox() ([]Box, error) {
	var box []Box
	o := orm.NewOrm()
	_, err := o.QueryTable("box").All(&box)
	return box, err
}

func FindBoxByName(name string) (Box, error) {
	var box Box
	o := orm.NewOrm()
	err := o.QueryTable("box").Filter("name", name).One(&box)
	return box, err
}

func FindBoxById(id int) (Box, error) {
	var box Box
	o := orm.NewOrm()
	err := o.QueryTable("box").Filter("id", id).One(&box)
	return box, err
}

type Boxinfo struct {
	Box
	Name    string `json:"name"`
	Address string `json:"address"`
	Series  string
	Model   string `json:"model";valid:"Required;MaxSize(20)"`
	Status  string `json:"status";valid:"Required"`
}

func PaginatorBox(page, size int, name string) ([]Boxinfo, int64, error) {
	o := orm.NewOrm()
	var boxinfo []Boxinfo
	sql := "select a.id,a.owner_id,a.boxid, a.installdate,a.serialno,a.service,a.zoomowner,a.hqowner,a.usernums,a.state,a.remark,b.name,b.address, c.model, c.status from box a left join owner b on a.owner_id=b.id  left join stock c on a.serialno=c.serialno"
	offset := size * (page - 1)
	if name == "" {
		count, _ := o.Raw(sql).QueryRows(&boxinfo)
		sql += " limit ?  offset ?"
		_, err := o.Raw(sql, size, offset).QueryRows(&boxinfo)
		return boxinfo, count, err
	} else {
		cname := "%" + name + "%"
		sql += " where a.boxid like ? or a.serialno like ?  or a.service like ?  or b.name like ? or c.model like ? "
		count, _ := o.Raw(sql, cname, cname, cname, cname, cname).QueryRows(&boxinfo)
		sql += " limit ? offset ?"
		_, err := o.Raw(sql, cname, cname, cname, cname, cname, size, offset).QueryRows(&boxinfo)
		return boxinfo, count, err
	}

}

func BoxExistCheck(name string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("box").Filter("boxid", name).Exist()
	return exist
}

func BoxAdd(boxinfo *Boxinfo) error {
	o := orm.NewOrm()
	err := o.Begin()
	StockStatusChange(boxinfo.Serialno, boxinfo.Status, o)
	AddStockout(boxinfo.Series, boxinfo.Model, boxinfo.Serialno, boxinfo.Status, boxinfo.Boxid, boxinfo.Name, boxinfo.Installdate, o)
	err = AddBox(&boxinfo.Box, o)
	if err != nil {
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}

func AddBox(box *Box, o orm.Ormer) error {
	box.State = "在线"
	_, err := o.Insert(box)
	return err
}

func BoxEdit(info *Boxinfo) error {
	o := orm.NewOrm()
	err := o.Begin()
	StockStatusChange(info.Serialno, info.Status, o)
	_, err = EditBox(info.Id, info.Owner_id, info.Installdate, info.Service, info.State, info.Usernums, info.Remark, o)
	if err != nil {
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}

func EditBox(id, owner_id int, installdate, service, state string, usernums int, remark string, o orm.Ormer) (int64, error) {
	sql := "update box set owner_id=?,installdate=?,service=?,state=?,usernums=?,remark=? where id=?"
	res, err := o.Raw(sql, owner_id, installdate, service, state, usernums, remark, id).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}

func DeleteBox(id int, o orm.Ormer) (int64, error) {
	if num, err := o.Delete(&Box{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

func BoxDelete(id, stockid, serialno int, indate, remark string) error {
	o := orm.NewOrm()
	err := o.Begin()
	StockStatusChange(serialno, "空闲机", o)
	AddStockin(stockid, indate, remark, o)
	_, err = DeleteBox(id, o)
	if err != nil {
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}
