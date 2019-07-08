package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) ApiV1Info() {
	c.Data["Result"] = "Api for something. Version 0.1"
	c.TplName = "ApiV1Info.tpl"
}
