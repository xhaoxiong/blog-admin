package models

type Axis struct {
	Id            int64
	Title         string `form:"Title"`
	Content       string `form:"Content"`
	Time          string `form:"Time"`
	Delete        int    `form:"Delete"`
	ContentDetail string `form:"ContentDetail";gorm:"type:longtext"`
}
