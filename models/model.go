package models

import (
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id      int    `json:"id"`
	Series  string `json:"series"`
	Model   string `json:"model"`
	Comment string `json:"comment"`
}

func init() {
	orm.RegisterModel(new(Model))
}

func FindAllModel() ([]Model, error) {
	var model []Model
	o := orm.NewOrm()
	_, err := o.QueryTable("model").All(&model)
	return model, err
}

func FindModelBySeries(series string) ([]Model, error) {
	var model []Model
	o := orm.NewOrm()
	_, err := o.QueryTable("model").Filter("series", series).All(&model)
	return model, err
}
