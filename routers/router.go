// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"officeoa/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/api/v1/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/api/v1/logout", &controllers.LoginController{}, "get:Logout")
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSNamespace("/owner",
				beego.NSRouter("/", &controllers.OwnerController{}),
				beego.NSRouter("/add", &controllers.OwnerController{}, "post:AddOwner"),
				beego.NSRouter("/customer", &controllers.OwnerController{}, "get:GetOwnername"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.OwnerController{}, "delete:Delete"),
				beego.NSRouter("/info/:id([0-9]+)", &controllers.OwnerController{}, "get:Ownerinfo"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.OwnerController{}, "post:Editowner"),
			),
			beego.NSNamespace("/contact",
				beego.NSRouter("/", &controllers.ContactController{}, "get:Getcontact"),
				beego.NSRouter("/add", &controllers.ContactController{}, "post:Add"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.ContactController{}, "post:Edit"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.ContactController{}, "delete:Delete"),
			),
			beego.NSNamespace("/box",
				beego.NSRouter("/", &controllers.BoxController{}),
				beego.NSRouter("/add", &controllers.BoxController{}, "post:Add"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.BoxController{}, "post:Edit"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.BoxController{}, "delete:Delete"),
			),
			beego.NSNamespace("/stock",
				beego.NSRouter("/", &controllers.StockController{}),
				beego.NSRouter("/add", &controllers.StockController{}, "post:Add"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.StockController{}, "post:Edit"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.StockController{}, "delete:Delete"),
			),
			beego.NSNamespace("/model",
				beego.NSRouter("/", &controllers.ModelController{}),
				beego.NSRouter("/add", &controllers.ModelController{}, "post:Add"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.ModelController{}, "post:Edit"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.ModelController{}, "delete:Delete"),
			),
			beego.NSNamespace("/series",
				beego.NSRouter("/", &controllers.SeriesController{}),
				beego.NSRouter("/add", &controllers.SeriesController{}, "post:Add"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.SeriesController{}, "post:Edit"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.SeriesController{}, "delete:Delete"),
			),
			beego.NSNamespace("/project",
				beego.NSRouter("/", &controllers.ProjectController{}),
				beego.NSRouter("/add", &controllers.ProjectController{}, "post:Add"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.ProjectController{}, "post:Edit"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.ProjectController{}, "delete:Delete"),
			),

			beego.NSNamespace("/users",
				beego.NSRouter("/", &controllers.UserController{}, "get:GetAll"),
				beego.NSRouter("/add", &controllers.UserController{}, "post:Post"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.UserController{}, "post:Edit"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.UserController{}, "delete:Delete"),
				beego.NSRouter("/profile/:id([0-9]+)", &controllers.UserController{}, "get:GetUserprofile"),
			),
			beego.NSNamespace("/permission",
				beego.NSRouter("/", &controllers.PermissionController{}, "get:Get"),
				beego.NSRouter("/all", &controllers.PermissionController{}, "get:GetAll"),
				beego.NSRouter("/add", &controllers.PermissionController{}, "post:Add"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.PermissionController{}, "post:Edit"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.PermissionController{}, "delete:Delete"),
			),
			beego.NSNamespace("/roles",
				beego.NSRouter("/", &controllers.RoleController{}, "get:Get"),
				beego.NSRouter("/add", &controllers.RoleController{}, "post:Add"),
				beego.NSRouter("/edit/:id([0-9]+)", &controllers.RoleController{}, "post:Edit"),
				beego.NSRouter("/delete/:id([0-9]+)", &controllers.RoleController{}, "delete:Delete"),
				beego.NSRouter("/prole/:id([0-9]+)", &controllers.RoleController{}, "get:PermissionRole"),
				beego.NSRouter("/addprole/:id([0-9]+)", &controllers.RoleController{}, "post:AddPermissionRole"),
				beego.NSRouter("/delprole/:id([0-9]+)", &controllers.RoleController{}, "post:DeletePermissionRole"),
				beego.NSRouter("/userrole", &controllers.RoleController{}, "get:UserRole"),
				beego.NSRouter("/edituserrole/:id([0-9]+)", &controllers.RoleController{}, "post:EditUserRole"),
			),
		),
	)
	beego.AddNamespace(ns)
}
