package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"officeoa/models"
	"strconv"
	"strings"
)

type StockController struct {
	FrontController
}

func (c *StockController) Get() {
	var page, size int
	var name string
	c.Ctx.Input.Bind(&page, "page")
	c.Ctx.Input.Bind(&size, "size")
	c.Ctx.Input.Bind(&name, "name")
	name = strings.TrimSpace(name)
	stock, total, err := models.PaginatorStock(page, size, name)
	c.checkerr(err, "get box error")
	result := make(map[string]interface{})
	result["stock"] = stock
	result["total"] = total
	c.Response(true, "success", result)
}

func (c *StockController) Add() {
	info := &models.Stocks{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	valid := validation.Validation{}
	valid.Required(info.Serialno, "serialno").Message("序列号不能空")
	valid.Range(info.Serialno, 1000, 99999999, "serialno").Message("序列号输入数字不对")
	valid.Required(info.Series, "series").Message("系列不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(false, "验证失败: "+err.Key+" , "+err.Message, "")
		}
	}
	if info.Serialno == 0 {
		c.Response(false, "序列号输入错误", "")
	}
	check := models.SerialnoExistCheck(info.Serialno)
	if check {
		c.Response(false, "序列号已经存在", "")
	}
	err := models.StockAdd(info)
	c.checkerr(err, "add data error")
	c.Response(true, "添加数据成功", "")
}

func (c *StockController) Edit() {
	id := c.Ctx.Input.Param(":id")
	sid, err := strconv.Atoi(id)
	c.checkerr(err, "id atoi error")
	info := &models.Stock{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	valid := validation.Validation{}
	valid.Required(info.Serialno, "serialno").Message("序列号不能空")
	valid.Range(info.Serialno, 1000, 99999999, "serialno").Message("序列号输入数字不对")
	valid.Required(info.Series, "series").Message("系列不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(false, "验证失败: "+err.Key+" , "+err.Message, "")
		}
	}
	if info.Serialno == 0 {
		c.Response(false, "序列号输入错误", "")
	}
	_, err = models.EditStock(info, sid)
	c.checkerr(err, "add data error")
	c.Response(true, "编辑数据成功", "")
}

func (c *StockController) Delete() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "id atoi error")
	err = models.StockDelete(bid)
	if err != nil {
		c.Response(false, "删除库存出错", err)
	}
	c.Response(true, "删除成功", "")
}
