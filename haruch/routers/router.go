package routers

import (
	"beegotest/haruch/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.PostController{})
}
