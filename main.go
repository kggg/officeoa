package main

import (
	_ "officeoa/routers"
	//"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	/*
		beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
			ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
		}))
	*/
	//beego.InsertFilter("/admin/", beego.BeforeRouter, controllers.Permissioncheck)
	beego.Run()
}
