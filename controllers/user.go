package controllers

import (
	"encoding/json"
	"officeoa/models"
	//"github.com/astaxie/beego"
	"github.com/gogather/com"
	//"log"
	"strconv"
	"strings"
)

// Operations about Users
type UserController struct {
	FrontController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.Users	true		"body for user content"
// @Success 200 {int} models.Users.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid, err := models.AddUser(user)
	u.checkerr(err, "add user error")
	u.Response(true, "success", uid)
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.Users
// @router / [get]
func (u *UserController) GetAll() {
	var page, size int
	var name string
	u.Ctx.Input.Bind(&page, "page")
	u.Ctx.Input.Bind(&size, "size")
	u.Ctx.Input.Bind(&name, "name")
	name = strings.TrimSpace(name)
	users, total, err := models.PaginatorUser(page, size, name)
	u.checkerr(err, "get all user info error")
	result := make(map[string]interface{})
	result["users"] = users
	result["total"] = total
	u.Response(true, "success", result)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	id := u.GetString(":id")
	if id != "" {
		iuid, err := strconv.Atoi(id)
		u.checkerr(err, "uid atoi error")
		user, uerr := models.FindUserById(iuid)
		if uerr != nil {
			u.Response(false, "get user info error", uerr.Error())
		} else {
			u.Response(true, "success", user)
		}
	} else {
		u.Response(false, "user id empty", "")
	}
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Users	true		"body for user content"
// @Success 200 {object} models.Users
// @Failure 403 :uid is not int
// @router /:uid [put]

func (u *UserController) Edit() {
	id := u.GetString(":id")
	if id != "" {
		var user models.User
		uid, err := strconv.Atoi(id)
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.EditUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":id")
	iuid, err := strconv.Atoi(uid)
	u.checkerr(err, "uid atoi error")
	models.DeleteUser(iuid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	var userinfo models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &userinfo)

	if userinfo.Username == "" || userinfo.Password == "" {
		u.Response(false, "empty", "empty")
	}
	user, err := models.FindUserByName(userinfo.Username)
	if err != nil {
		u.Response(false, "username does not exists", userinfo.Username)

	} else {
		pass := com.Md5(userinfo.Password)
		if pass == user.Password {

			u.SetSession("username", user.Username)
			u.SetSession("userid", user.Id)
			u.Response(true, "success", "")
		} else {
			u.Response(false, "passwd invalid", userinfo.Password)
		}
	}

}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.SetSession("username", nil)
	u.SetSession("userid", nil)
	u.DelSession("username")
	u.DelSession("userid")
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// get user profile
func (c *UserController) GetUserprofile() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "id atoi error")
	up, err := models.FindUserprofileById(bid)
	c.checkerr(err, "get user profile error")
	c.Response(true, "success", up)
}
