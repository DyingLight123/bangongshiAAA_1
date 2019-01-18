package routers

import (
	"bangongshiAAA_1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/influxdb", &controllers.InfluxdbController{})

	beego.Router("/refresh", &controllers.RefreshController{})
}
