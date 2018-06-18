package controllers

import (
	"blogadmin/models"
	"io/ioutil"
	"strconv"
	"log"
)

// @router /api/admin/article/list [*]
func (this *AdminController) ListArticle() {

	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	per, _ := this.GetInt("per")
	if per == 0 {
		per = 10
	}

	qs := models.DB.Model(models.Article{})

	count := 0
	qs.Count(&count)
	var articles []*models.Article
	qs.Limit(per).Offset((page - 1) * per).Order("id desc").Find(&articles)

	this.ReturnSuccess("articles", articles, "page", page, "count", count, "per", per)
	return
}

// @router /api/admin/article/detail [*]
func (this *AdminController) ArticleDetail() {
	var article models.Article

	if _, err := this.GetByID(&article); err != nil {
		this.ReturnJson(10001, "article not found")
		return
	}

	data, err := ioutil.ReadFile("/var/www/server/static/article" + strconv.Itoa(int(article.Id)) + ".txt")

	if err != nil {
		log.Println("read article file error")
		this.ReturnJson(10002, "read article file error")
		return
	}

	article.Content = string(data)

	this.ReturnSuccess("article", article)
}

// @router /api/admin/article/save [*]
func (this *AdminController) SaveArticle() {

	var article models.Article
	if id, err := this.GetByID(&article); err != nil {
		if id != 0 {
			this.ReturnJson(10001, "article not found")
			return
		}

	}

	if err := this.ParseForm(&article); err != nil {
		this.ReturnJson(10002, "parse article error")
		return
	}

	if err := models.DB.Save(&article).Error; err != nil {
		this.ReturnJson(10003, "article save error")
		return
	}

	err := ioutil.WriteFile("/var/www/server/static/article"+strconv.Itoa(int(article.Id))+".txt", []byte(article.Content), 0777)

	if err != nil {
		log.Println("write file error:", err)
		this.ReturnJson(10003, "write file error")
		return
	}
	this.ReturnSuccess()
}

// @router /api/admin/article/delete [*]
func (this *AdminController) DeleteArticle() {
	var article models.Article
	if _, err := this.GetByID(&article); err != nil {
		this.ReturnJson(10001, "article not found")
		return
	}

	article.Delete = 1

	if err := models.DB.Save(&article).Error; err != nil {
		this.ReturnJson(10002, "delete error")
		return
	}
	this.ReturnSuccess()

}
