package models

import (
	"github.com/astaxie/beego/orm"
)

type Userprofile struct {
	Id         int
	User_id    int    `json:"user_id"`
	Department string `json:"department"`
	Name       string `json:"name"`
	Birthday   string `json:"birthday"`
	Mobile     string `json:"mobile"`
	Telephone  string `json:"telephone"`
	Email      string `json:"email"`
	Qq         int64  `json:"qq"`
	Address    string `json:"address"`
	Workdate   string `json:"workdate"`
}

func init() {
	orm.RegisterModel(new(Userprofile))
}

func FindUserprofileById(uid int) (Userprofile, error) {
	o := orm.NewOrm()
	var up Userprofile
	err := o.QueryTable("userprofile").Filter("user_id", uid).One(&up)
	return up, err
}
