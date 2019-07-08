package routers

import (
	"beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/v1", &controllers.MainController{}, "get:ApiV1Info")
}
