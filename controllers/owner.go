package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"officeoa/models"
	//"log"
	"strconv"
	"strings"
)

type OwnerController struct {
	FrontController
}

func (c *OwnerController) Get() {
	var page, size int
	var name string
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&size, "size")
	c.Ctx.Input.Bind(&name, "name")
	name = strings.TrimSpace(name)
	result := make(map[string]interface{})
	owners, total, err := models.PaginatorOwner(page, size, name)
	c.checkerr(err, "get hostinfo error")
	result["total"] = total
	result["data"] = owners

	c.Response(true, "success", result)

}

func (c *OwnerController) AddOwner() {
	info := &models.Owner{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	check := models.OwnerExistCheck(info.Name)
	if check {
		c.Response(false, "客户名称已经存在", "")
	}
	valid := validation.Validation{}
	valid.Required(info.Name, "name").Message("客户名称不能空")
	valid.Required(info.Address, "address").Message("地址不能为空")
	valid.Required(info.Type, "type").Message("类型不能为空")
	valid.MaxSize(info.Remark, 300, "remark").Message("字符不能超过300个")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(false, "输入错误", err.Key+" : "+err.Message)
		}
	}
	num, err := models.AddOwner(info.Name, info.Address, info.Type, info.Remark)
	c.checkerr(err, "addowner error")
	c.Response(true, "success", num)
}

func (c *OwnerController) Editowner() {
	info := &models.Owner{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	valid := validation.Validation{}
	valid.Required(info.Name, "name").Message("客户名称不能空")
	valid.Required(info.Address, "address").Message("地址不能为空")
	valid.Required(info.Type, "type").Message("类型不能为空")
	valid.MaxSize(info.Remark, 300, "remark").Message("字符不能超过300个")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(false, "输入错误", err.Key+" : "+err.Message)
		}
	}
	num, err := models.EditOwner(info.Id, info.Name, info.Address, info.Type, info.Remark)
	c.checkerr(err, "editowner error")
	c.Response(true, "success", num)
}

func (c *OwnerController) Delete() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "id atoi error")
	rid, err := models.DeleteOwner(bid)
	if err != nil {
		c.Response(false, "删除客户出错", err)
	}
	c.Response(true, "删除成功", rid)
}

func (c *OwnerController) Ownerinfo() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "id atoi error")
	info, err := models.FindContactById(bid)
	c.checkerr(err, "find ownerinfo by id error")
	c.Response(true, "success", info)
}

func (c *OwnerController) GetOwnername() {
	ownername, err := models.FindAllOwner()
	c.checkerr(err, "get owner name error")
	c.Response(true, "success", ownername)
}
