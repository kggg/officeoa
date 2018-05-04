package models

import (
	"github.com/astaxie/beego/orm"
)

type Contact struct {
	Id       int
	Cname    string `valid:"Required; MaxSize(50)"`
	Owner_id int    `valid:"Required; Min(1);Max(90000)"`
	Position string `valid:"MaxSize(30)"`
	Email    string `valid:"Email; MaxSize(100)"`
	Phone1   string `valid:"Required; MaxSize(50)"`
	Phone2   string `valid:"Numeric; MaxSize(50)"`
}

func init() {
	orm.RegisterModel(new(Contact))
}

type Ownercontact struct {
	Contact
	Name string
}

func FindAllContact() ([]Ownercontact, error) {
	var info []Ownercontact
	o := orm.NewOrm()
	sql := "select a.cname, a.email,a.phone1,a.phone2,b.name from contact a left join owner b on a.owner_id=b.id"
	_, err := o.Raw(sql).QueryRows(&info)
	return info, err
}

func FindContactById(id int) ([]Ownercontact, error) {
	var info []Ownercontact
	o := orm.NewOrm()
	sql := "select a.cname, a.position,a.email,a.phone1,a.phone2,b.name from contact a left join owner b on a.owner_id=b.id where a.owner_id=?"
	_, err := o.Raw(sql, id).QueryRows(&info)
	return info, err
}

func PaginatorContact(page, size int, name string) ([]Ownercontact, int64, error) {
	o := orm.NewOrm()
	var contact []Ownercontact
	sql := "select a.id,a.owner_id,a.cname, a.position,a.email,a.phone1,a.phone2,b.name from contact a left join owner b on a.owner_id=b.id"
	offset := size * (page - 1)
	if name == "" {
		count, _ := o.Raw(sql).QueryRows(&contact)
		sql += " limit ?  offset ?"
		_, err := o.Raw(sql, size, offset).QueryRows(&contact)
		return contact, count, err
	} else {
		cname := "%" + name + "%"
		sql += " where a.cname like ? or a.email like ?  or a.phone1 like ?  or a.phone2 like ? or b.name like ? "
		count, _ := o.Raw(sql, cname, cname, cname, cname, cname).QueryRows(&contact)
		sql += " limit ? offset ?"
		_, err := o.Raw(sql, cname, cname, cname, cname, cname, size, offset).QueryRows(&contact)
		return contact, count, err
	}

}

func AddContact(contact *Contact) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(contact)
	return id, err
}

func EditContact(contact *Contact) (int64, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("contact").Filter("id", contact.Id).Update(orm.Params{
		"cname":    contact.Cname,
		"owner_id": contact.Owner_id,
		"position": contact.Position,
		"email":    contact.Email,
		"phone1":   contact.Phone1,
		"phone2":   contact.Phone2,
	})
	return num, err
}

func DeleteContact(id int) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Contact{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}
