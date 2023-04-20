package controllers

import (
	"tfoms_server/classes/models"

	beego "github.com/beego/beego/v2/server/web"
)

type JournalController struct {
	beego.Controller
}

func (c *JournalController) Get() {
	c.Data["Website"] = "tf.site"
	c.Data["Email"] = "john-bursa3@mail.ru"
	c.Data["EmailName"] = "Neverhood"
	filters := make(map[string]string)
	filters["organization_id"] = "1"
	patients, _ := models.GetPatientJournal(filters, 1000)
	c.Data["patients"] = patients
	c.TplName = "index.tpl"
}
