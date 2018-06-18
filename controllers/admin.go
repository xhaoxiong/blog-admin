package controllers

import (
	"github.com/astaxie/beego"
	"blogadmin/models"
)

type Common struct {
	beego.Controller
	User    models.AdminUser
	IsLogin bool
}

func (this *Common) GetByID(obj interface{}) (int64, error) {
	id, _ := this.GetInt64("id")
	return id, models.DB.Where("id=?", id).First(obj).Error
}

type AdminController struct {
	Common
}

type AdminCommonController struct {
	Common
}

func (this *AdminController) Prepare() {
	AdminInterceptor(&this.Common)
}

func (this *Common) ReturnJson(status int, message string, args ...interface{}) {
	result := make(map[string]interface{})
	result["status"] = status
	result["message"] = message

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}

	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

func (this *Common) ReturnSuccess(args ...interface{}) {
	result := make(map[string]interface{})
	result["status"] = 10000
	result["message"] = "success"
	key := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}
	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

//@router /admin/* [*]
func (this *AdminCommonController) AdminIndex() {
	this.TplName = "admin/index.html"
}
