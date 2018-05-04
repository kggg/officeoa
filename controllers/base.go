package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Response(status bool, str string, data interface{}) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str, "data": data}
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) checkerr(err error, str string) {
	if err != nil {
		this.Response(false, str, err)
	}
}

func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
