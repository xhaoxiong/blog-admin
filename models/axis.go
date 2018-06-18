package models

type Axis struct {
	Id            int64
	Title         string `form:"title"`
	Content       string `form:"content"`
	Time          string `form:"time"`
	Delete        int    `form:"delete"`
	ContentDetail string `form:"content_detail";gorm:"type:longtext"`
}
