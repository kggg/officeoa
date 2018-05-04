package models

import (
	"github.com/astaxie/beego/orm"
)

type UserRole struct {
	User_id int `orm:"pk"; json:"user_id"`
	Role_id int `json:"role_id"`
}

func init() {
	orm.RegisterModel(new(UserRole))
}

type Urole struct {
	Id       int    `json:"id"`
	Rname    string `json:"rname"`
	Username string `json:"username"`
}

func FindAllUserRole() ([]Urole, error) {
	o := orm.NewOrm()
	var ur []Urole
	sql := " SELECT T0.`user_id`,T0.`role_id`,T2.`rname`,T2.`username` FROM `user_role` T0 INNER JOIN `role` T1 ON T1.`id` = T0.`role_id` INNER JOIN `user` T2 ON T2.`id` = T0.`user_id`"
	_, err := o.Raw(sql).QueryRows(&ur)
	return ur, err
}

func PaginatorUserRole(page, size int, name string) ([]Urole, int64, error) {
	o := orm.NewOrm()
	var ur []Urole
	sql := " SELECT T0.`id`,T2.`rname`,T0.`username` FROM `user` T0 LEFT JOIN `user_role` T1 ON T1.`user_id` = T0.`id` LEFT JOIN `role` T2 ON T2.`id` = T1.`role_id`"
	offset := size * (page - 1)
	if name == "" {
		count, _ := o.Raw(sql).QueryRows(&ur)
		sql += " limit ?  offset ?"
		_, err := o.Raw(sql, size, offset).QueryRows(&ur)
		return ur, count, err
	} else {
		cname := "%" + name + "%"
		sql += " where T2.rname like ? or T0.username like ? "
		count, _ := o.Raw(sql, cname, cname).QueryRows(&ur)
		sql += " limit ? offset ?"
		_, err := o.Raw(sql, cname, cname, size, offset).QueryRows(&ur)
		return ur, count, err
	}

}

func FindUserRoleByUserId(uid int) (Urole, error) {
	o := orm.NewOrm()
	var ur Urole
	sql := " SELECT T1.`id`, T1.`rname`,T2.`username` FROM `user_role` T0 INNER JOIN `role` T1 ON T1.`id` = T0.`role_id` INNER JOIN `user` T2 ON T2.`id` = T0.`user_id` WHERE T0.`user_id` = ? "
	err := o.Raw(sql, uid).QueryRow(&ur)
	return ur, err
}

func DeleteUserRole(rid, uid int) (int64, error) {
	o := orm.NewOrm()
	sql := "delete from user_role where role_id=? and user_id=?"
	res, err := o.Raw(sql, rid, uid).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}

func AddUserRole(rid, uid int) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into user_role (role_id, user_id) values(?,?)"
	res, err := o.Raw(sql, rid, uid).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}

func EditUserRole(rid, uid int) (int64, error) {
	o := orm.NewOrm()
	sql := "update user_role set role_id=? where user_id=?"
	res, err := o.Raw(sql, rid, uid).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}

func UserRoleExistCheck(uid int) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("user_role").Filter("user_id", uid).Exist()
	return exist
}
