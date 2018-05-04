package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"officeoa/models"
	//"github.com/astaxie/beego"
	//"log"
	"strconv"
	"strings"
	"time"
)

type RoleController struct {
	FrontController
}

func (c *RoleController) Get() {
	var page, size int
	var name string
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&size, "size")
	c.Ctx.Input.Bind(&name, "name")
	name = strings.TrimSpace(name)
	roles, total, err := models.PaginatorRole(page, size, name)
	c.checkerr(err, "get all user info error")
	result := make(map[string]interface{})
	result["data"] = roles
	result["total"] = total
	c.Response(true, "success", result)
}

func (c *RoleController) Add() {
	info := &models.Role{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	valid := validation.Validation{}
	valid.Required(info.Rname, "name").Message("权限角色名不能空")
	valid.MaxSize(info.Description, 50, "description").Message("不能超过50个字")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(false, "验证失败: "+err.Key+" , "+err.Message, "")
		}
	}
	check := models.RoleExistCheck(info.Rname)
	if check {
		c.Response(false, "权限角色已经存在", "")
	}
	info.Created_at = time.Now().Format("2006-01-02 15:04:05")
	_, err := models.AddRole(info)
	c.checkerr(err, "add data error")
	c.Response(true, "添加权限角色成功", "")
}

func (c *RoleController) Edit() {
	info := &models.Role{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	valid := validation.Validation{}
	valid.Required(info.Rname, "name").Message("权限角色不能空")
	valid.MaxSize(info.Description, 50, "description").Message("不能超过50个字")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(false, "验证失败: "+err.Key+" , "+err.Message, "")
		}
	}
	_, err := models.EditRole(info)
	c.checkerr(err, "add data error")
	c.Response(true, "添加权限名成功", "")
}

func (c *RoleController) Delete() {
	pid := c.GetString(":id")
	iuid, err := strconv.Atoi(pid)
	c.checkerr(err, "id atoi error")
	_, err = models.DeleteRole(iuid)
	if err != nil {
		c.Response(false, "数据库删除该权限角色失败", err.Error())
	}
	c.Response(true, "success!", "")
}

func (c *RoleController) PermissionRole() {
	rid := c.GetString(":id")
	iid, err := strconv.Atoi(rid)
	c.checkerr(err, "id atoi error")
	permissionrole, err := models.FindPermissionRoleByRoleId(iid)
	c.checkerr(err, "get permission role info error")
	c.Response(true, "success", permissionrole)
}

func (c *RoleController) AddPermissionRole() {
	rid := c.GetString(":id")
	iid, err := strconv.Atoi(rid)
	c.checkerr(err, "id atoi error")
	var info []int
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	if len(info) < 1 {
		c.Response(false, "没有选中权限", "")
	} else if len(info) == 1 {
		_, err := models.AddPermissionRole(info[0], iid)
		c.checkerr(err, "add prole error")
	} else {
		for _, v := range info {
			_, err := models.AddPermissionRole(v, iid)
			c.checkerr(err, "add prole error")
		}
	}

	c.Response(true, "success", iid)
}

func (c *RoleController) DeletePermissionRole() {
	rid := c.GetString(":id")
	iid, err := strconv.Atoi(rid)
	c.checkerr(err, "id atoi error")
	var info []int
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	if len(info) < 1 {
		c.Response(false, "没有选中权限", "")
	} else if len(info) == 1 {
		_, err := models.DeletePermissionRole(info[0], iid)
		c.checkerr(err, "Delete prole error")
	} else {
		for _, v := range info {
			_, err := models.DeletePermissionRole(v, iid)
			c.checkerr(err, "Delete prole error")
		}
	}

	c.Response(true, "success", iid)
}

func (c *RoleController) UserRole() {
	var page, size int
	var name string
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&size, "size")
	c.Ctx.Input.Bind(&name, "name")
	name = strings.TrimSpace(name)
	roles, total, err := models.PaginatorUserRole(page, size, name)
	c.checkerr(err, "get all user role info error")
	result := make(map[string]interface{})
	result["data"] = roles
	result["total"] = total
	c.Response(true, "success", result)
}

func (c *RoleController) EditUserRole() {
	id := c.GetString(":id")
	uid, err := strconv.Atoi(id)
	c.checkerr(err, "uid atoi error")
	type Urole struct {
		Id       int
		Rname    int
		Username string
	}
	info := &Urole{}
	json.Unmarshal(c.Ctx.Input.RequestBody, info)
	if info.Id == 0 || info.Rname == 0 {
		c.Response(false, "提交数据出错", "")
	}
	var num int64
	check := models.UserRoleExistCheck(uid)
	if check {
		num, err = models.EditUserRole(info.Rname, info.Id)
		c.checkerr(err, "database error")
	} else {
		num, err = models.AddUserRole(info.Rname, info.Id)
		c.checkerr(err, "database error")
	}
	c.Response(true, "success", num)
}
