package controllers

import (

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"blogadmin/models"
)

func CaptchaInterceptor(ctr *Common) {
	code := ctr.GetString("captcha_code")
	captchaId := ctr.GetString("captcha_id")
	if !captcha.VerifyString(captchaId, code) {
		ctr.ReturnJson(10401, "captcha verify error")
		ctr.StopRun()
		return
	}
}

func AdminInterceptor(ctr *Common) {
	if userinfo := ctr.GetSession("admininfo"); userinfo != nil {
		user := userinfo.(models.AdminUser)
		if err := models.DB.Where("id=? ", user.ID).First(&user).Error; err != nil {
			beego.Warning("user read error:", err)
			ctr.Abort("403")
			ctr.StopRun()
			return
		}

		ctr.User = user
		ctr.IsLogin = true

		return
	}
	beego.Debug("[UserIntercepor] user not found")
	ctr.Abort("403")
	ctr.StopRun()

}
