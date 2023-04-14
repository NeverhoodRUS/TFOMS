package controllers

import (
	"tfoms_server/classes/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tf.site"
	c.Data["Email"] = "john-bursa3@mail.ru"
	c.Data["EmailName"] = "Neverhood"
	filters := make(map[string]string)
	filters["organization_id"] = "1"
	patients, _ := models.GetPatientJournal(filters)
	c.Data["patients"] = patients
	c.TplName = "jc.tpl"
}
