package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"] = append(beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"] = append(beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"] = append(beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"] = append(beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/getall`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"] = append(beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/post`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"] = append(beego.GlobalControllerRouter["beegotest/haruch/controllers:PostController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/post/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
