package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tf.site"
	c.Data["Email"] = "john-bursa3@mail.ru"
	c.Data["EmailName"] = "Neverhood"
	c.TplName = "index.tpl"
}
