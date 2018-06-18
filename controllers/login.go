package controllers

import (

	"time"
	"blogadmin/models"
)

// @router /api/admin/logout [*]
func (this *AdminCommonController) Logout() {

	this.DelSession("admininfo")
	this.ReturnSuccess()

}

// @router /api/admin/login [*]
func (this *AdminCommonController) AdminLogin() {
	//CaptchaInterceptor(&this.Common)
	username := this.GetString("username")
	password := this.GetString("password")

	var user models.AdminUser

	if err := models.DB.
		Where("username=? and password = ?", username,password).
		First(&user).Error; err != nil {
		this.ReturnJson(10001, "username or password not found")
		return
	}
	models.DB.Model(&user).Updates(map[string]interface{}{"last_login_ip": this.Ctx.Input.IP(), "last_login_time": time.Now()})
	this.SetSession("admininfo", user)
	user.Password = ""
	this.ReturnSuccess("user", user)
}
