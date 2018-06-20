package models

type Article struct {
	Id         int64
	Title      string `form:"Title"`
	Name       string `form:"Name"`
	Ps         string `form:"Ps"`
	Text       string `form:"Text"`
	Date       string `form:"Date"`
	Reading    int    `form:"Reading"`
	Agree      int    `form:"Agree"`
	Delete     int    `form:"Delete"`
	Exhibition int    `form:"Exhibition"`
	Type       int    `form:"Type"`
	Content    string `form:"Content";gorm:"type:longtext"`
}
