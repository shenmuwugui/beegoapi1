// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beegoapi/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/home",
			beego.NSInclude(
				&controllers.HomeController{},
			),
		),

		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),

		beego.NSNamespace("/college",
			beego.NSInclude(
				&controllers.CollegeController{},
			),
		),

		beego.NSNamespace("/deforment",
			beego.NSInclude(
				&controllers.DeformentController{},
			),
		),
	)
	beego.AddNamespace(ns)

	//beego.Router("/del/:ids", &controllers.StudentController{}, "get:DelMany")
	beego.Router("/deldeforment/:ids", &controllers.DeformentController{}, "patch:DelMany")
}
