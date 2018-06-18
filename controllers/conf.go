package controllers

import (
	"blogadmin/models"
)

// @router /api/admin/conf/save [*]
func (this *AdminController) SaveConf() {

	var conf models.Conf
	if id, err := this.GetByID(&conf); err != nil {
		if id != 0 {
			this.ReturnJson(10001, "conf not found")
			return
		}

	}

	if err := this.ParseForm(&conf); err != nil {
		this.ReturnJson(10002, "parse conf error")
		return
	}

	if err := models.DB.Save(&conf).Error; err != nil {
		this.ReturnJson(10003, "conf save error")
		return
	}

	this.ReturnSuccess()
}
