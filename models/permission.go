package models

import (
	"github.com/astaxie/beego/orm"
)

type Permission struct {
	Id          int    `json:"id"`
	Name        string `json:"name"; valid:"Required;MaxSize(15)`
	Path        string `json:"path"; valid:"Required;MaxSize(20)"`
	Description string `json:"description"; valid:"MaxSize(50)"`
	Created_at  string `json:"created_at"`
}

func init() {
	orm.RegisterModel(new(Permission))
}

func FindAllPermission() ([]Permission, error) {
	o := orm.NewOrm()
	var per []Permission
	_, err := o.QueryTable("permission").OrderBy("path").All(&per)
	return per, err
}

func PaginatorPermission(page, size int, name string) ([]Permission, int64, error) {
	o := orm.NewOrm()
	var per []Permission
	if name == "" {
		qs := o.QueryTable("permission")
		count, _ := qs.Count()
		qs = qs.OrderBy("path")

		_, err := qs.Limit(size).Offset(size * (page - 1)).All(&per)
		return per, count, err
	} else {
		cond := orm.NewCondition()
		cond1 := cond.And("name__contains", name).Or("path__icontains", name).Or("description__contains", name)
		qs := o.QueryTable("permission")
		qs = qs.SetCond(cond1)
		count, _ := qs.Count()
		qs = qs.OrderBy("path")
		qs = qs.Limit(size).Offset(size * (page - 1))
		_, err := qs.All(&per)
		return per, count, err
	}
}

func PermissionExistCheck(name string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("permission").Filter("name", name).Exist()
	return exist
}

func AddPermission(p *Permission) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(p)
	return id, err
}

func EditPermission(p *Permission) (int64, error) {
	o := orm.NewOrm()
	per := Permission{Id: p.Id}
	err := o.Read(&per)
	if err == nil {
		num, err := o.Update(p)

		return num, err
	}
	return 0, err
}

func DeletePermission(id int) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Delete(&Permission{Id: id})
	return num, err
}
