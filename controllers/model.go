package controllers

import (
	"officeoa/models"
	//"encoding/json"
	//"github.com/astaxie/beego/validation"
	//"log"
	//"strings"
)

type ModelController struct {
	FrontController
}

type Options struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Children []*Children
}

type Children struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

func (c *ModelController) Get() {
	var series string
	c.Ctx.Input.Bind(&series, "series")
	model, err := models.FindModelBySeries(series)
	c.checkerr(err, "find all model error")
	c.Response(true, "success", model)
}

func (c *ModelController) Add() {
}

func (c *ModelController) Edit() {
}
func (c *ModelController) Delete() {
}
