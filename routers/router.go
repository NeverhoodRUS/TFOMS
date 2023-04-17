package routers

import (
	"tfoms_server/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/journal", &controllers.JournalController{})
}
