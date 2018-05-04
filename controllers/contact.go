package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"officeoa/models"
	//"log"
	"strconv"
	"strings"
)

type ContactController struct {
	FrontController
}

func (c *ContactController) Add() {
	var ownerinfo = &models.Contact{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ownerinfo)
	valid := validation.Validation{}
	b, err := valid.Valid(ownerinfo)
	if err != nil {
		c.Response(false, "do validate error", err.Error())
	}
	if !b {
		for _, err := range valid.Errors {
			c.Response(false, "validation error: ", err.Key+"- "+err.Message)
		}
	}
	inertid, err := models.AddContact(ownerinfo)
	c.checkerr(err, "add contact info error")
	c.Response(true, "success", inertid)
}

func (c *ContactController) Edit() {
	var ownerinfo = &models.Contact{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ownerinfo)
	valid := validation.Validation{}
	b, err := valid.Valid(ownerinfo)
	if err != nil {
		c.Response(false, "do validate error", err.Error())
	}
	if !b {
		for _, err := range valid.Errors {
			c.Response(false, "validation error: ", err.Key+"- "+err.Message)
		}
	}
	inertid, err := models.EditContact(ownerinfo)
	c.checkerr(err, "add contact info error")
	c.Response(true, "success", inertid)
}

func (c *ContactController) Getcontact() {
	var page, size int
	var name string
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&size, "size")
	c.Ctx.Input.Bind(&name, "name")
	name = strings.TrimSpace(name)
	result := make(map[string]interface{})
	contact, total, err := models.PaginatorContact(page, size, name)
	c.checkerr(err, "get contact list error")
	result["total"] = total
	result["contact"] = contact
	result["status"] = true
	result["info"] = "success"
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}

func (c *ContactController) Delete() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "id atoi error")
	rid, err := models.DeleteContact(bid)
	if err != nil {
		c.Response(false, "删除联系人出错", err)
	}
	c.Response(true, "删除成功", rid)
}
