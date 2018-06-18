package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"blogadmin/controllers"
)

func init() {
	beego.Include(&controllers.AdminController{},
		&controllers.AdminCommonController{})

	beego.Router("/api/ueditor_controller", &controllers.Ueditor{}, "*:U_Controller")
	beego.SetStaticPath("/upload", "upload")
	beego.Handler("/api/image/captcha/*.png", captcha.Server(90, 40))
}
