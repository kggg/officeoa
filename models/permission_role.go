package models

import (
	"github.com/astaxie/beego/orm"
)

type PermissionRole struct {
	Permission *Permission `orm:"rel(fk)"; json:"permission_id"`
	Role       *Role       `orm:"pk;rel(fk)"; json:"role_id"`
}

func init() {
	orm.RegisterModel(new(PermissionRole))
}

type Prole struct {
	Permission
	Rname string `json:"rname"`
}

func FindPermissionRoleByRoleId(rid int) ([]Prole, error) {
	o := orm.NewOrm()
	var pr []Prole
	sql := " SELECT T1.`id`, T1.`name`, T1.`path`,T2.`rname` FROM `permission_role` T0 INNER JOIN `permission` T1 ON T1.`id` = T0.`permission_id` INNER JOIN `role` T2 ON T2.`id` = T0.`role_id` WHERE T0.`role_id` = ? "
	_, err := o.Raw(sql, rid).QueryRows(&pr)
	return pr, err
}

func DeletePermissionRole(pid, rid int) (int64, error) {
	o := orm.NewOrm()
	sql := "delete from permission_role where permission_id=? and role_id=?"
	res, err := o.Raw(sql, pid, rid).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}

func AddPermissionRole(pid, rid int) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into permission_role (permission_id, role_id) values(?,?)"
	res, err := o.Raw(sql, pid, rid).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}
