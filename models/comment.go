package models

type Comment struct {
	Id            int64
	Name          string `form:"Name"`
	Time          string `form:"Time"`
	Content       string `form:"Content"`
	Delete        int    `form:"Delete"`
	ContentDetail string `form:"ContentDetail";gorm:"type:longtext"`
}
