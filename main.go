package main

import (
	_ "moshopserver/controllers"
	_ "moshopserver/models"
	_ "moshopserver/routers"
	"moshopserver/services"
	_ "moshopserver/services"
	_ "moshopserver/utils"

	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 8080

	beego.InsertFilter("/api/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8080

}
