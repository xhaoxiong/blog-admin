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

// @router /api/admin/conf/list [*]
func (this *AdminController) ListConf() {

	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	per, _ := this.GetInt("per")
	if per == 0 {
		per = 10
	}

	qs := models.DB.Model(models.Conf{})

	count := 0
	qs.Count(&count)
	var confs []*models.Conf
	qs.Limit(per).Offset((page - 1) * per).Order("id desc").Find(&confs)

	this.ReturnSuccess("confs", confs, "page", page, "count", count, "per", per)
	return
}
