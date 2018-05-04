package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Status     int    `json:"status"`
	Created_at string `json:"created_at"`
	Update_at  string
}

type Userinfo struct {
	User
	Userprofile
}

func init() {
	orm.RegisterModel(new(User))
}

func PaginatorUser(page, size int, name string) ([]User, int64, error) {
	o := orm.NewOrm()
	var userinfo []User
	sql := "select id, username, status, created_at from user"
	offset := size * (page - 1)
	if name == "" {
		count, _ := o.Raw(sql).QueryRows(&userinfo)
		sql += " limit ?  offset ?"
		_, err := o.Raw(sql, size, offset).QueryRows(&userinfo)
		return userinfo, count, err
	} else {
		cname := "%" + name + "%"
		sql += " where username like ?  "
		count, _ := o.Raw(sql, cname).QueryRows(&userinfo)
		sql += " limit ? offset ?"
		_, err := o.Raw(sql, cname, size, offset).QueryRows(&userinfo)
		return userinfo, count, err
	}

}

func FindAllUser() ([]User, error) {
	var user []User
	o := orm.NewOrm()
	_, err := o.QueryTable("user").All(&user)
	return user, err
}

func FindUserByName(name string) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("username", name).Filter("status", 1).One(&user)
	return user, err
}

func FindUserById(id int) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("id", id).One(&user)
	return user, err
}

func AddUser(u User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&u)
	return id, err
}

func EditUser(id int, u *User) (int64, error) {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err == nil {
		num, err := o.Update(u)

		return num, err
	}
	return 0, err
}

func DeleteUser(id int) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Delete(&User{Id: id})
	return num, err

}

func UserExistCheck(name string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("user").Filter("name", name).Exist()
	return exist
}

func FindUserPermissionById(uid int) ([]string, error) {
	o := orm.NewOrm()
	var path []string
	sql := "select f.path from user a inner join user_role b on a.id=b.user_id inner join role c on b.role_id=c.id inner join permission_role d on d.role_id=c.id inner join permission f on f.id=d.permission_id where a.id=?"
	_, err := o.Raw(sql, uid).QueryRows(&path)
	return path, err
}

func FindUserRoleById(uid int) (string, error) {
	o := orm.NewOrm()
	var role string
	sql := "select c.rname from user a inner join user_role b on a.id=b.user_id inner join role c on b.role_id=c.id  where a.id=?"
	err := o.Raw(sql, uid).QueryRow(&role)
	return role, err

}
