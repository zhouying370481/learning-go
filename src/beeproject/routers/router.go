package routers

import (
	"github.com/pathbox/learning-go/src/beeproject/routers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
