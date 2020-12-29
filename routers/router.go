package routers

import (
	"Branches/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/main", &controllers.MainController{})
    beego.Router("/web" )
}
