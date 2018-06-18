package controllers

import "github.com/dchest/captcha"

// @router /api/common/captcha [*]
func (this *AdminCommonController) GetCaptcha() {
	captchaId := captcha.NewLen(4)
	result := make(map[string]interface{})
	result["status"] = 10000
	result["src"] = "/api/image/captcha/" + captchaId + ".png"
	result["id"] = captchaId
	this.Data["json"] = result
	this.ServeJSON()
	return
}
