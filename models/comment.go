package models

type Comment struct {
	Id            int64
	Name          string `form:"name"`
	Time          string `form:"time"`
	Content       string `form:"content"`
	Delete        int    `form:"delete"`
	ContentDetail string `form:"content_detail";gorm:"type:longtext"`
}
