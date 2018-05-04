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

// Operations about Users
type PermissionController struct {
	FrontController
}

func (c *PermissionController) Get() {
	var page, size int
	var name string
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&size, "size")
	c.Ctx.Input.Bind(&name, "name")
	name = strings.TrimSpace(name)
	permission, total, err := models.PaginatorPermission(page, size, name)
	c.checkerr(err, "get all user info error")
	result := make(map[string]interface{})
	result["permission"] = permission
	result["total"] = total
	c.Response(true, "success", result)
}

func (c *PermissionController) GetAll() {
	permission, err := models.FindAllPermission()
	c.checkerr(err, "get all permission info error")
	c.Response(true, "success", permission)
}

func (c *PermissionController) Add() {
	info := &models.Permission{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	valid := validation.Validation{}
	valid.Required(info.Name, "name").Message("权限名不能空")
	valid.Required(info.Path, "path").Message("权限路径不能为空")
	valid.MaxSize(info.Path, 50, "path").Message("权限路径不能超过50个字符")
	valid.MaxSize(info.Description, 50, "description").Message("不能超过50个字")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(false, "验证失败: "+err.Key+" , "+err.Message, "")
		}
	}
	check := models.PermissionExistCheck(info.Name)
	if check {
		c.Response(false, "权限名已经存在", "")
	}
	info.Created_at = time.Now().Format("2006-01-02 15:04:05")
	_, err := models.AddPermission(info)
	c.checkerr(err, "add data error")
	c.Response(true, "添加权限名成功", "")
}

func (c *PermissionController) Edit() {
	info := &models.Permission{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	valid := validation.Validation{}
	valid.Required(info.Name, "name").Message("权限名不能空")
	valid.Required(info.Path, "path").Message("权限路径不能为空")
	valid.MaxSize(info.Path, 50, "path").Message("权限路径不能超过50个字符")
	valid.MaxSize(info.Description, 50, "description").Message("不能超过50个字")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(false, "验证失败: "+err.Key+" , "+err.Message, "")
		}
	}
	_, err := models.EditPermission(info)
	c.checkerr(err, "add data error")
	c.Response(true, "添加权限名成功", "")
}

func (c *PermissionController) Delete() {
	pid := c.GetString(":id")
	iuid, err := strconv.Atoi(pid)
	c.checkerr(err, "uid atoi error")
	_, err = models.DeletePermission(iuid)
	if err != nil {
		c.Response(false, "数据库删除该权限失败", err.Error())
	}
	c.Response(true, "success!", "")
}
