package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"officeoa/models"
	//"log"
	"strconv"
	"strings"
)

type BoxController struct {
	FrontController
}

func (c *BoxController) Get() {
	var page, size int
	var name string
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&size, "size")
	c.Ctx.Input.Bind(&name, "name")
	name = strings.TrimSpace(name)
	box, total, err := models.PaginatorBox(page, size, name)
	c.checkerr(err, "get box error")
	result := make(map[string]interface{})
	result["box"] = box
	result["total"] = total
	c.Response(true, "success", result)

}

// validation 验证还没有生效
func (c *BoxController) Add() {
	info := &models.Boxinfo{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	valid := validation.Validation{}
	b, err := valid.Valid(info)
	if err != nil {
		c.Response(false, "do validate error", err.Error())
	}
	if !b {
		for _, err := range valid.Errors {
			c.Response(false, "validation error: ", err.Key+"- "+err.Message)
		}
	}
	if info.Serialno == 0 {
		c.Response(false, "序列号输入错误", "")
	}

	check := models.BoxExistCheck(info.Boxid)
	if check {
		c.Response(false, "设备boxid名称已经存在", "")
	}
	stock, err := models.FindStockBySerialno(info.Serialno)
	if err != nil {
		c.Response(false, "输入的序列号不存在", "")
	}
	if stock.Status != "空闲机" {
		c.Response(false, "该序列号已经被其它设备使用中", "")
	}

	err = models.BoxAdd(info)
	if err != nil {
		c.Response(false, "新增设备失败", err.Error())
	}
	c.Response(true, "新增设备成功", "")

}

func (c *BoxController) Edit() {
	info := &models.Boxinfo{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	valid := validation.Validation{}
	valid.Required(info.Owner_id, "owner_id").Message("客户名称不能空")
	valid.Required(info.State, "state").Message("设备状态不能为空")
	valid.Required(info.Service, "service").Message("服务类型不能为空")
	valid.Required(info.Status, "status").Message("购买类型不能为空")
	valid.MaxSize(info.Remark, 300, "remark").Message("字符不能超过300个")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(false, "输入错误", err.Key+" : "+err.Message)
		}
	}
	err := models.BoxEdit(info)
	c.checkerr(err, "editowner error")
	c.Response(true, "success", "")
}

func (c *BoxController) Delete() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "id atoi error")
	var company string
	c.Ctx.Input.Bind(&company, "company")
	boxinfo, err := models.FindBoxById(bid)
	c.checkerr(err, "find this id from box table error")
	stock, err := models.FindStockBySerialno(boxinfo.Serialno)
	c.checkerr(err, "datebase error")
	remark := company + " ; " + boxinfo.Boxid + " 停止使用"
	err = models.BoxDelete(boxinfo.Id, stock.Id, boxinfo.Serialno, boxinfo.Installdate, remark)
	if err != nil {
		c.Response(false, "删除设备出错", err.Error())
	}
	c.Response(true, "删除成功", "")
}
