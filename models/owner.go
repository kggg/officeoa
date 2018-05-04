package models

import (
	"github.com/astaxie/beego/orm"
)

type Owner struct {
	Id         int
	Name       string
	Address    string
	Type       string
	Remark     string
	Created_at string
}

func init() {
	orm.RegisterModel(new(Owner))
}

func FindAllOwner() ([]Owner, error) {
	var owner []Owner
	o := orm.NewOrm()
	_, err := o.QueryTable("owner").All(&owner)
	return owner, err
}

func FindOwnerAllName() ([]orm.ParamsList, error) {
	o := orm.NewOrm()
	var list []orm.ParamsList
	_, err := o.QueryTable("owner").ValuesList(&list, "id", "name")
	return list, err
}

func PaginatorOwner(page, size int, name string) ([]Owner, int64, error) {
	o := orm.NewOrm()
	var owner []Owner
	if name == "" {
		qs := o.QueryTable("owner")
		count, _ := qs.Count()

		_, err := qs.Limit(size).Offset(size * (page - 1)).All(&owner)
		return owner, count, err
	} else {
		cond := orm.NewCondition()
		cond1 := cond.And("name__icontains", name).Or("address__icontains", name)
		qs := o.QueryTable("owner")
		qs = qs.SetCond(cond1)
		count, _ := qs.Count()
		qs = qs.Limit(size).Offset(size * (page - 1))
		_, err := qs.All(&owner)
		return owner, count, err
	}

}

func OwnerCount() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("owner").Filter("type", "终端用户").Count()
	return cnt, err
}

func FindOwnerByName(name string) (Owner, error) {
	var owner Owner
	o := orm.NewOrm()
	err := o.QueryTable("owner").Filter("name", name).One(&owner)
	return owner, err
}

func FindOwnerById(id int) (Owner, error) {
	var owner Owner
	o := orm.NewOrm()
	err := o.QueryTable("owner").Filter("id", id).One(&owner)
	return owner, err
}

func AddOwner(name, address, pty, remark string) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into owner (name, address , type, remark) values(?, ?, ?, ?)"
	res, err := o.Raw(sql, name, address, pty, remark).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func OwnerExistCheck(name string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("owner").Filter("name", name).Exist()
	return exist
}

func DeleteOwner(id int) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Owner{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

func EditOwner(id int, name, address, pty, remark string) (int64, error) {
	o := orm.NewOrm()
	sql := "update owner set name=?, address=? , type=?, remark=? where id=?"
	res, err := o.Raw(sql, name, address, pty, remark, id).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}
