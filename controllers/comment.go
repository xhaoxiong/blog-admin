package controllers

import (
	"io/ioutil"
	"strconv"
	"blogadmin/models"
	"log"
)

// @router /api/admin/comment/list [*]
func (this *AdminController) ListComment() {

	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	per, _ := this.GetInt("per")
	if per == 0 {
		per = 10
	}

	qs := models.DB.Model(models.Comment{})

	count := 0
	qs.Count(&count)
	var comments []*models.Comment
	qs.Limit(per).Offset((page - 1) * per).Order("id desc").Find(&comments)

	this.ReturnSuccess("comments", comments, "page", page, "count", count, "per", per)
	return
}

// @router /api/admin/comment/detail [*]
func (this *AdminController) CommentsDetail() {
	var comment models.Comment

	if _, err := this.GetByID(&comment); err != nil {
		this.ReturnJson(10001, "comment not found")
		return
	}

	data, err := ioutil.ReadFile("/var/www/server/static/comment" + strconv.Itoa(int(comment.Id)) + ".txt")

	if err != nil {
		log.Println("read article file error")
		this.ReturnJson(10002, "read article file error")
		return
	}

	comment.ContentDetail = string(data)

	this.ReturnSuccess("comment", comment)
}

// @router /api/admin/article/delete [*]
func (this *AdminController) DeleteComment() {
	var comment models.Comment
	if _, err := this.GetByID(&comment); err != nil {
		this.ReturnJson(10001, "article not found")
		return
	}

	comment.Delete = 1

	if err := models.DB.Save(&comment).Error; err != nil {
		this.ReturnJson(10002, "delete error")
		return
	}
	this.ReturnSuccess()

}
