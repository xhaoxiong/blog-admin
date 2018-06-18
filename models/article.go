package models

type Article struct {
	Id         int64
	Title      string `form:"title"`
	Name       string `form:"name"`
	Ps         string `form:"ps"`
	Text       string `form:"text"`
	Date       string `form:"date"`
	Reading    int    `form:"reading"`
	Agree      int    `form:"agree"`
	Delete     int    `form:"delete"`
	Exhibition int    `form:"exhibition"`
	Type       int    `form:"type"`
	Content    string `form:"content";gorm:"type:longtext"`
}
