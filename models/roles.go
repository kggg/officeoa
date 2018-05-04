package models

import (
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id          int    `json:"id"`
	Rname       string `json:"rname"; valid:"Required;MaxSize(15)`
	Description string `json:"description"; valid:"MaxSize(50)"`
	Created_at  string `json:"created_at"`
}

func init() {
	orm.RegisterModel(new(Role))
}

func PaginatorRole(page, size int, name string) ([]Role, int64, error) {
	o := orm.NewOrm()
	var role []Role
	if name == "" {
		qs := o.QueryTable("role")
		count, _ := qs.Count()
		_, err := qs.Limit(size).Offset(size * (page - 1)).All(&role)
		return role, count, err
	} else {
		cond := orm.NewCondition()
		cond1 := cond.And("rname__contains", name).Or("description__contains", name)
		qs := o.QueryTable("role")
		qs = qs.SetCond(cond1)
		count, _ := qs.Count()
		qs = qs.Limit(size).Offset(size * (page - 1))
		_, err := qs.All(&role)
		return role, count, err
	}
}

func RoleExistCheck(name string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("role").Filter("rname", name).Exist()
	return exist
}

func AddRole(r *Role) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(r)
	return id, err
}

func EditRole(r *Role) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: r.Id}
	err := o.Read(&role)
	if err == nil {
		num, err := o.Update(r)

		return num, err
	}
	return 0, err
}

func DeleteRole(id int) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Delete(&Role{Id: id})
	return num, err
}
