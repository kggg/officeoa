package models

import (
	"github.com/astaxie/beego/orm"
)

type Modelseries struct {
	Id     int    `json:"id"`
	Series string `json:"series"`
}

func init() {
	orm.RegisterModel(new(Modelseries))
}

func FindSeries() ([]Modelseries, error) {
	o := orm.NewOrm()
	var series []Modelseries
	_, err := o.QueryTable("modelseries").OrderBy("series").All(&series)
	return series, err
}
