package models

import (
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id       int    `json:"id"`
	Date     string `json:"date"`
	Pname    string `json:"pname"`
	Creater  string `json:"creater"`
	Engineer string `json:"engineer"`
	Model    string `json:"model"`
	Serialno int    `json:"serialno"`
	Boxid    string `json:"boxid"`
	Customer string `json:"customer"`
	Status   string `json:"status"`
	Comments string `json:"comments"`
	Updated  string `json:"updated"`
}

func init() {
	orm.RegisterModel(new(Project))
}

func FindAllProject() ([]Project, error) {
	o := orm.NewOrm()
	var project []Project
	_, err := o.QueryTable("project").All(&project)
	return project, err
}

func FindProjectById(id int) (Project, error) {
	o := orm.NewOrm()
	var project Project
	err := o.QueryTable("project").Filter("id", id).One(&project)
	return project, err
}

func PaginatorProject(page, size int, name string) ([]Project, int64, error) {
	o := orm.NewOrm()
	var project []Project
	if name == "" {
		qs := o.QueryTable("project")
		count, _ := qs.Count()
		qs = qs.OrderBy("-id")
		_, err := qs.Limit(size).Offset(size * (page - 1)).All(&project)
		return project, count, err
	} else {
		cond := orm.NewCondition()
		cond1 := cond.And("pname__contains", name).Or("model__icontains", name).Or("boxid__icontains", name).Or("status__contains", name).Or("engineer__contains", name).Or("customer__contains", name).Or("comments__contains", name).Or("serialno__contains", name)
		qs := o.QueryTable("project")
		qs = qs.SetCond(cond1)
		count, _ := qs.Count()
		qs = qs.OrderBy("-id")
		qs = qs.Limit(size).Offset(size * (page - 1))
		_, err := qs.All(&project)
		return project, count, err
	}
}

func AddProject(pro *Project) (int64, error) {
	o := orm.NewOrm()
	pro.Status = "未分配"
	num, err := o.Insert(pro)
	return num, err
}

func DeleteProject(id int) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Project{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}

}
