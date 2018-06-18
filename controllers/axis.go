package controllers

import (
	"io/ioutil"
	"strconv"
	"blogadmin/models"
	"log"
)

// @router /api/admin/axis/list [*]
func (this *AdminController) ListAxis() {

	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	per, _ := this.GetInt("per")
	if per == 0 {
		per = 10
	}

	qs := models.DB.Model(models.Axis{})

	count := 0
	qs.Count(&count)
	var Axises []*models.Axis
	qs.Limit(per).Offset((page - 1) * per).Order("id desc").Find(&Axises)

	this.ReturnSuccess("axises", Axises, "page", page, "count", count, "per", per)
	return
}

// @router /api/admin/axis/detail [*]
func (this *AdminController) AxisDetail() {
	var axis models.Axis

	if _, err := this.GetByID(&axis); err != nil {
		this.ReturnJson(10001, "axis not found")
		return
	}

	data, err := ioutil.ReadFile("/var/www/server/static/axis" + strconv.Itoa(int(axis.Id)) + ".txt")

	if err != nil {
		log.Println("read article file error")
		this.ReturnJson(10002, "read article file error")
		return
	}

	axis.Content = string(data)

	this.ReturnSuccess("axis", axis)
}

// @router /api/admin/axis/save [*]
func (this *AdminController) SaveAxis() {

	var axis models.Axis
	if id, err := this.GetByID(&axis); err != nil {
		if id != 0 {
			this.ReturnJson(10001, "axis not found")
			return
		}

	}

	if err := this.ParseForm(&axis); err != nil {
		this.ReturnJson(10002, "parse axis error")
		return
	}

	if err := models.DB.Save(&axis).Error; err != nil {
		this.ReturnJson(10003, "axis save error")
		return
	}

	err := ioutil.WriteFile("/var/www/server/static/ais"+strconv.Itoa(int(axis.Id))+".txt", []byte(axis.Content), 0777)

	if err != nil {
		log.Println("write file error:", err)
		this.ReturnJson(10003, "write file error")
		return
	}
	this.ReturnSuccess()
}

// @router /api/admin/axis/delete [*]
func (this *AdminController) DeleteAxis() {
	var axis models.Axis
	if _, err := this.GetByID(&axis); err != nil {
		this.ReturnJson(10001, "article not found")
		return
	}

	axis.Delete = 1

	if err := models.DB.Save(&axis).Error; err != nil {
		this.ReturnJson(10002, "delete error")
		return
	}
	this.ReturnSuccess()

}
