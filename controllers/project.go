package controllers

import (
	"officeoa/models"
	//"encoding/json"
	//"github.com/astaxie/beego/validation"
	//"log"
	"strconv"
	"strings"
)

type ProjectController struct {
	FrontController
}

func (c *ProjectController) Get() {
	var page, size int
	var name string
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&size, "size")
	c.Ctx.Input.Bind(&name, "name")
	name = strings.TrimSpace(name)
	project, total, err := models.PaginatorProject(page, size, name)
	c.checkerr(err, "get project info error")
	result := make(map[string]interface{})
	result["box"] = project
	result["total"] = total
	c.Response(true, "success", result)
}

func (c *ProjectController) Add() {
}
func (c *ProjectController) Edit() {
}

func (c *ProjectController) Delete() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "id atoi error")
	rid, err := models.DeleteProject(bid)
	if err != nil {
		c.Response(false, "删除项目出错", err)
	}
	c.Response(true, "删除成功", rid)
}
