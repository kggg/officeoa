package controllers

import (
	"officeoa/models"
	//"log"
	"regexp"
	"strings"
)

type FrontController struct {
	BaseController
}

func (c *FrontController) Prepare() {
	username := c.GetSession("username")
	if username == nil {
		c.Response(false, "还没有登录", 400011)
	}
	/*
	 */
	userid := c.GetSession("userid").(int)
	role, err := models.FindUserRoleById(userid)
	c.checkerr(err, "get user role error")
	if role != "超级管理员" {
		permission := c.Permissioncheck(userid)
		if !permission {
			c.Response(false, "你没有权限访问", 400022)
		}
	}
}

func (c *FrontController) Permissioncheck(uid int) bool {
	paths, err := models.FindUserPermissionById(uid)
	c.checkerr(err, "find user permission error")
	if len(paths) == 0 {
		return false
	}
	path := c.Ctx.Request.RequestURI
	path = strings.Replace(path, "/api/v1", "", 1)
	path = strings.Split(path, "?")[0]
	for _, p := range paths {
		if path != "/" {
			if p == "/" {
				continue
			}
		}
		pindex := strings.LastIndex(path, "/")
		findex := strings.LastIndex(p, "/")
		if path[0:pindex] != "" && p[0:findex] == "" {
			continue
		}
		if path[0:pindex] == "" && p[0:findex] != "" {
			continue
		}
		if path[0:pindex] != "" && p[0:findex] != "" && p[0:findex] != path[0:pindex] {
			continue
		}
		if p[findex+1:] == path[pindex+1:] {
			return true
		}
		if p[findex+1:] == "" {
			continue
		}
		reg, err := regexp.Compile(p[findex+1:])
		if err != nil {
			continue
		}
		if reg.MatchString(path[pindex+1:]) {
			return true
		}

	}
	return false
}
