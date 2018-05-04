package controllers

import (
	"encoding/json"
	"github.com/gogather/com"
	"officeoa/models"
	//"log"
)

type LoginController struct {
	BaseController
}

func (u *LoginController) Login() {
	var userinfo models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &userinfo)

	if userinfo.Username == "" || userinfo.Password == "" {
		u.Response(false, "用户名和密码不能为空", "empty")
	}
	user, err := models.FindUserByName(userinfo.Username)
	if err != nil {
		u.Response(false, "用户名不存在", userinfo.Username)

	} else {
		pass := com.Md5(userinfo.Password)
		if pass == user.Password {

			u.SetSession("username", user.Username)
			u.SetSession("userid", user.Id)
			u.Response(true, "success", user.Username)
		} else {
			u.Response(false, "密码错误", userinfo.Password)
		}
	}

}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *LoginController) Logout() {
	u.SetSession("username", nil)
	u.DelSession("username")
	u.SetSession("userid", nil)
	u.DelSession("userid")
	u.Response(true, "logout success", 400012)
}
