package main

import (
	_ "tfoms_server/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
